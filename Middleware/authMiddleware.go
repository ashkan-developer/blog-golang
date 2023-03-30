package Middleware

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"mymodule/Database"
	"mymodule/Helper"
	"mymodule/Model"
	"net/http"
	"time"
)

var userCollection *mongo.Collection = Database.OpentCollection(Database.Client, "user")

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		clientToken := c.Request().Header.Get("token")
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		//var user Model.User
		var foundUser Model.User

		if clientToken == "" {
			res := c.JSON(http.StatusInternalServerError, "No Authorization header provided")
			return res
		}

		claims, err := Helper.ValidateToken(clientToken)

		user_find := userCollection.FindOne(ctx, bson.M{"email": claims.Email}).Decode(&foundUser)

		if foundUser.User_type == "0" {
			res := c.JSON(http.StatusInternalServerError, "error")
			return res
		}

		if user_find != nil {
			res := c.JSON(http.StatusInternalServerError, "error in the find user")
			return res
		}

		if err != "" {
			res := c.JSON(http.StatusInternalServerError, "error")
			return res
		}
		defer cancel()
		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("uid", claims.Uid)
		return next(c)
	}

}
