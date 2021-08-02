package videoCollection

import (
	"V1/models"
	"V1/repository"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const collectionName = "videos"

type VideoCollection interface {
	Aggregate(ctx context.Context, pipeline interface{}) (*models.Video, error)
	InsertOne(ctx context.Context, user interface{}) (primitive.ObjectID, error)
}
type videoCollection struct {
	db repository.DatabaseHelper
}

func newVideoCollection(db repository.DatabaseHelper) VideoCollection {
	return &videoCollection{
		db: db,
	}
}

func (v *videoCollection) Aggregate(ctx context.Context, pipeline interface{}) (*models.Video, error) {
	result, err := v.db.Collection(collectionName).Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	var video []*models.Video
	err = result.All(ctx, &video)
	if err != nil {
		return nil, err
	}
	return video[0], nil
}

func (v *videoCollection) InsertOne(ctx context.Context, user interface{}) (primitive.ObjectID, error) {
	newUser, err := v.db.Collection(collectionName).InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return newUser.InsertedID.(primitive.ObjectID), nil
}
