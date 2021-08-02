package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

//Video is struct for entity of video data from user
type Video struct {
	Title       string             `json:"title,omitempty" binding:"require"`
	Description string             `json:"description,omitempty" binding:"require"`
	URL         string             `json:"url,omitempty" binding:"require"`
	Author      primitive.ObjectID `json:"author" binding:"require"`
}
