package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" binding:"required"`
	Firstname string             `bson:"firstname,omitempty" json:"firstname,omitempty" binding:"required"`
	Lastname  string             `bson:"lastname,omitempty" json:"lastname,omitempty" binding:"required" `
	Age       int8               `bson:"age,omitempty" json:"age,omitempty" binding:"gte=1,lte=130,required"`
	RefVideo  []Video            `bson:"refVideo,omitempty" json:"ref,omitempty" `
}
