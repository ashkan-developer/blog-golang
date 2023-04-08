package Controller

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"log"
	"mymodule/Database"
	"mymodule/Helper"
	"mymodule/Model"
	"net/http"
	"os"
	"time"
)

var blogCollection *mongo.Collection = Database.OpentCollection(Database.Client, "blog")

func Blog_All(c echo.Context) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	cursor, err := blogCollection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}
	var episodes []bson.M
	if err = cursor.All(ctx, &episodes); err != nil {
		log.Fatal(err)
	}
	defer cancel()

	return c.JSON(http.StatusOK, episodes)
}
func Create_Blog(c echo.Context) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var blog Model.Blog

	if err := c.Bind(&blog); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	count, err := userCollection.CountDocuments(ctx, bson.M{"title": blog.Title})
	defer cancel()

	if err != nil {
		log.Panic(err)
		resp := c.JSON(http.StatusBadRequest, Helper.ErroLog(http.StatusInternalServerError, "error occured while checking for the title", "EXT_REF"))
		return resp
	}

	if count > 0 {
		resp := c.JSON(http.StatusInternalServerError, Helper.ErroLog(http.StatusInternalServerError, " this title already exists", "EXT_REF"))
		return resp
	}

	//file, err := c.FormFile("image")
	file, err := c.FormFile("image")

	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("Image/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	blog.Image = file
	resultInsertionNumber, insertErr := blogCollection.InsertOne(ctx, blog)
	if insertErr != nil {
		fmt.Sprintf("User item was not created")
		resp := c.JSON(http.StatusBadRequest, Helper.ErroLog(http.StatusInternalServerError, "error", "EXT_REF"))
		return resp
	}
	defer cancel()
	return c.JSON(http.StatusOK, resultInsertionNumber)
}

func Blog(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	var blog Model.Blog
	blog_name := c.Param("title")

	cursor := blogCollection.FindOne(ctx, bson.M{"title": blog_name}).Decode(&blog)

	if cursor != nil {
		log.Fatal("err")
	}

	defer cancel()
	return c.JSON(http.StatusOK, blog)
}

func DestroyBlog(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	blog_name := c.Param("title")

	res, err := blogCollection.DeleteOne(ctx, bson.M{"title": blog_name})

	if err != nil {
		log.Fatal("DeleteOne() ERROR:", err)
	}
	if res.DeletedCount == 0 {
		fmt.Println("DeleteOne() document not found:", res)
	}

	defer cancel()

	return c.JSON(http.StatusOK, res)
}

func UpdateBlog(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	//var category Model.Category
	blog_name := c.Param("title")
	title_new := c.FormValue("title")
	introduction_new := c.FormValue("introduction")
	category_name_new := c.FormValue("category_name")

	//		file		//

	file, err_file := c.FormFile("image")

	if err_file != nil {
		log.Fatal("Error file:", err_file)
	}
	src, err := file.Open()

	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("Image/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	filter := bson.M{"title": blog_name}

	// create the update query
	update := bson.D{
		{"$set",
			bson.D{
				{"title", title_new},
				{"image", file},
				{"introduction", introduction_new},
				{"category_name", category_name_new},
			},
		},
	}

	res, err := blogCollection.UpdateMany(ctx, filter, update)

	if err != nil {
		log.Fatal("DeleteOne() ERROR:", err)
	}
	if res.UpsertedCount == 0 {
		fmt.Println("DeleteOne() document not found:", res)
	}

	defer cancel()

	return c.JSON(http.StatusOK, res)
}
