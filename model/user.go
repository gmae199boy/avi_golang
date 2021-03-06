package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name string `json:"name,omitempty"`
	Email string `json:"email,omitempty" binding:"required"`
	Password string `json:"password,omitempty"`
}