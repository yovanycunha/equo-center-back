package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID 			primitive.ObjectID 	`json:"id" bson:"_id, omitempty"`
	Email		string `json:"email" binding:"required"`
	Password 	string `json:"password" binding:"required"`
}