package models

import "github.com/mongodb/mongo-go-driver/bson/primitive"

type Article struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Title string             `json:"title"`
}
