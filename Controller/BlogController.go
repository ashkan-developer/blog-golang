package Controller

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"mymodule/Database"
	"net/http"
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
