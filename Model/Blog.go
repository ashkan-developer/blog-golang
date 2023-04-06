package Model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mime/multipart"
)

type Blog struct {
	ID           primitive.ObjectID    `bson:"Id"`
	Title        *string               `bson:"title" json:"title" form:"title" validate:"required,min=2,max=100"`
	CategoryName *string               `bson:"category_name" json:"category_name" form:"category_name" validate:"required"`
	Image        *multipart.FileHeader `json:"image" bson:"image" form:"image" validate:"required"`
	Introduction *string               `json:"introduction" bson:"introduction" form:"introduction" validate:"required"`
}
