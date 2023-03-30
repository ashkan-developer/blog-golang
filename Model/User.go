package Model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID            primitive.ObjectID `bson:"id"`
	First_name    *string            `json:"first_name" form:"first_name" validate:"required,min=2,max=100"`
	Last_name     *string            `json:"last_name" form:"last_name" validate:"required,min=2,max=100"`
	Password      *string            `json:"Password" form:"Password" validate:"required,min=6"`
	Email         *string            `json:"email" form:"email" validate:"email,required"`
	Phone         *string            `json:"phone" form:"phone" validate:"required"`
	Token         *string            `json:"token"`
	User_type     string             `json:"user_type" bson:"User_type"`
	Refresh_token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id"`
}

type Error struct {
	ResponseCode      int    `json:"rc"`
	Message           string `json:"message"`
	Detail            string `json:"detail"`
	ExternalReference string `json:"ext_ref"`
}
