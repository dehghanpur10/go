package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//User is model that fetch from database
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Firstname string             `bson:"firstname,omitempty"`
	Lastname  string             `bson:"lastname,omitempty"`
	Age       int8               `bson:"age,omitempty"`
	RefVideo  []Video            `bson:"refVideo,omitempty"`
}
