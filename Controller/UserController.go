package Controller

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"log"
	"mymodule/Database"
	"mymodule/Helper"
	"mymodule/Model"
	"net/http"
	"time"
)

var userCollection *mongo.Collection = Database.OpentCollection(Database.Client, "user")
var validate = validator.New()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func Register(c echo.Context) error {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user Model.User

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	validationErr := validate.Struct(user)
	if validationErr != nil {
		resp := c.JSON(http.StatusBadRequest, Helper.ErroLog(http.StatusBadRequest, "erro", "EXT_REF"))
		return resp
	}

	count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	defer cancel()

	if err != nil {
		log.Panic(err)
		resp := c.JSON(http.StatusBadRequest, Helper.ErroLog(http.StatusInternalServerError, "error occured while checking for the email", "EXT_REF"))
		return resp
	}

	password := HashPassword(*user.Password)
	user.Password = &password

	if count > 0 {
		resp := c.JSON(http.StatusInternalServerError, Helper.ErroLog(http.StatusInternalServerError, " this email or phone number already exists", "EXT_REF"))
		return resp
	}
	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()
	token, refreshToken, _ := Helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, user.User_id)
	user.Token = &token
	user.Refresh_token = &refreshToken
	resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
	if insertErr != nil {
		fmt.Sprintf("User item was not created")
		resp := c.JSON(http.StatusBadRequest, Helper.ErroLog(http.StatusInternalServerError, "error", "EXT_REF"))
		return resp
	}
	defer cancel()
	return c.JSON(http.StatusOK, resultInsertionNumber)

}
