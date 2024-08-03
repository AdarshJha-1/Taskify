package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Title       string             `json:"title,omitempty" bson:"title"`
	IsCompleted bool               `json:"is_completed,omitempty" bson:"is_completed"`
	UserID      primitive.ObjectID `json:"user_id" bson:"user_id"`
}
