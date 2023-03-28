package Model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id          primitive.ObjectID `json:"id" bson:"id"`
	UserName    string             `json:"username" form:"username"`
	Email       string             `json:"email" form:"email"`
	Password    string             `json:"password" form:"password"`
	Token       string             `json:"token" form:"token"`
	DateCreated time.Time          `json:"date_created" form:"date_created"`
}
