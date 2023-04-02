package Model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mime/multipart"
)

type Category struct {
	ID    primitive.ObjectID    `bson:"id"`
	Title *string               `json:"title" bson:"title" form:"title" validate:"required,min=2,max=100"`
	Image *multipart.FileHeader `json:"image" bson:"image" form:"image" validate:"required"`
}
