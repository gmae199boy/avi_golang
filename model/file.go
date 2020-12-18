package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json: "_id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Path string `json:"path"`
	UserId primitive.ObjectID `bson:"userId" json:"userId"`
}