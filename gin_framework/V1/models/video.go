package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Video struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" binding:"required"`
	Title       string             `bson:"title,omitempty" json:"title,omitempty" binding:"required" `
	Description string             `bson:"description,omitempty" json:"description,omitempty" binding:"required"`
	URL         string             `bson:"url,omitempty" json:"url,omitempty" binding:"required" `
	Author      User               `bson:"author" json:"author" binding:"required"`
}
