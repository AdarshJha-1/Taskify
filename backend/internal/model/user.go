package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Username string             `json:"username,omitempty" bson:"username"`
	Email    string             `json:"email,omitempty" bson:"email"`
	Password string             `json:"password,omitempty" bson:"password"`
}
