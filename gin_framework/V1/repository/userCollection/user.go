package userCollection

import (
	"V1/models"
	"V1/repository"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const collectionName = "users"

type UserCollection interface {
	Aggregate(ctx context.Context, pipeline interface{}) (*models.User, error)
	InsertOne(ctx context.Context, user interface{}) (primitive.ObjectID, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}) error
}
type userDatabase struct {
	db repository.DatabaseHelper
}

func NewUserDatabase(db repository.DatabaseHelper) UserCollection {
	return &userDatabase{
		db: db,
	}
}

func (u *userDatabase) Aggregate(ctx context.Context, pipeline interface{}) (*models.User, error) {
	result, err := u.db.Collection(collectionName).Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	var user []*models.User
	err = result.All(ctx, &user)
	if err != nil {
		return nil, err
	}
	if len(user) == 0 {
		return nil, errors.New("user not found")
	}
	return user[0], nil
}

func (u *userDatabase) InsertOne(ctx context.Context, user interface{}) (primitive.ObjectID, error) {

	newUser, err := u.db.Collection(collectionName).InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return newUser.InsertedID.(primitive.ObjectID), nil
}
func (u *userDatabase) UpdateOne(ctx context.Context, filter interface{}, update interface{}) error {
	_, err := u.db.Collection(collectionName).UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
