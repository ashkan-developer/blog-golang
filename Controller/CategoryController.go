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

var categoryCollection *mongo.Collection = Database.OpentCollection(Database.Client, "category")

func Category_All(c echo.Context) error {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	cursor, err := categoryCollection.Find(ctx, bson.M{})

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

func Create_category(c echo.Context) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var category Model.Category

	if err := c.Bind(&category); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	count, err := userCollection.CountDocuments(ctx, bson.M{"title": category.Title})
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

	category.Image = file
	resultInsertionNumber, insertErr := categoryCollection.InsertOne(ctx, category)
	if insertErr != nil {
		fmt.Sprintf("User item was not created")
		resp := c.JSON(http.StatusBadRequest, Helper.ErroLog(http.StatusInternalServerError, "error", "EXT_REF"))
		return resp
	}
	defer cancel()
	return c.JSON(http.StatusOK, resultInsertionNumber)

}

func Category(c echo.Context) error {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	var category Model.Category
	category_name := c.Param("title")

	cursor := categoryCollection.FindOne(ctx, bson.M{"title": category_name}).Decode(&category)

	if cursor != nil {
		log.Fatal("err")
	}

	defer cancel()
	return c.JSON(http.StatusOK, category)

}

func DestroyCategory(c echo.Context) error {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	category_name := c.Param("title")

	res, err := categoryCollection.DeleteOne(ctx, bson.M{"title": category_name})

	if err != nil {
		log.Fatal("DeleteOne() ERROR:", err)
	}
	if res.DeletedCount == 0 {
		fmt.Println("DeleteOne() document not found:", res)
	}

	defer cancel()

	return c.JSON(http.StatusOK, res)

}

func UpdateCategory(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	//var category Model.Category
	category_name := c.Param("title")
	title_new := c.FormValue("title")

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

	filter := bson.M{"title": category_name}

	// create the update query
	update := bson.D{
		{"$set",
			bson.D{
				{"title", title_new},
				{"image", file},
			},
		},
	}

	res, err := categoryCollection.UpdateMany(ctx, filter, update)

	if err != nil {
		log.Fatal("DeleteOne() ERROR:", err)
	}
	if res.UpsertedCount == 0 {
		fmt.Println("DeleteOne() document not found:", res)
	}

	defer cancel()

	return c.JSON(http.StatusOK, res)
}
