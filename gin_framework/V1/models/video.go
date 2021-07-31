package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Video is model that fetch from data base
type Video struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	URL         string             `bson:"url,omitempty"`
	Author      User               `bson:"author,omitempty"`
}
