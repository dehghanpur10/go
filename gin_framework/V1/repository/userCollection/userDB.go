package userCollection

//
//import (
//	"V1/models"
//	"V1/repository"
//	"context"
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	"go.mongodb.org/mongo-driver/mongo"
//)
//
//const collectionName = "users"
//
//type UserDatabase interface {
//	FindOne(ctx context.Context, id string) (*models.User, error)
//	Create(ctx context.Context, usr *models.User) (interface{}, error)
//	Update(ctx context.Context, user string, video string) error
//}
//
//type userDatabase struct {
//	db repository.DatabaseHelper
//}
//
//func NewUserDatabase(db repository.DatabaseHelper) UserDatabase {
//	return &userDatabase{
//		db: db,
//	}
//}
//
//func (u *userDatabase) FindOne(ctx context.Context, id string) (*models.User, error) {
//	objectId, err := primitive.ObjectIDFromHex(id)
//	if err != nil {
//		return nil, err
//	}
//	filter := bson.D{{"$match", bson.D{{"_id", objectId}}}}
//	lookup := bson.D{{ // for lookup reference
//		"$lookup",
//		bson.D{
//			{"from", "videos"},
//			{"localField", "refVideo"},
//			{"foreignField", "_id"},
//			{"as", "refVideo"},
//		},
//	}}
//	pipeline := mongo.Pipeline{filter, lookup}
//	aggregate, err := u.db.Collection(collectionName).Aggregate(context.TODO(), pipeline)
//	if err != nil {
//		return nil, err
//	}
//	var users []models.User
//	err = aggregate.All(ctx, &users)
//	if err != nil {
//		return nil, err
//	}
//	return &users[0], nil
//}
//
//func (u *userDatabase) Create(ctx context.Context, usr *models.User) (interface{}, error) {
//	newUser, err := u.db.Collection(collectionName).InsertOne(ctx, usr)
//	if err != nil {
//		return nil, err
//	}
//	return newUser.InsertedID, nil
//}
//func (u *userDatabase) Update(ctx context.Context, user string, video string) error {
//	userID, err := primitive.ObjectIDFromHex(user)
//	if err != nil {
//		return nil
//	}
//	videoId, err := primitive.ObjectIDFromHex(video)
//	if err != nil {
//		return nil
//	}
//	_, err = u.db.Collection(collectionName).UpdateOne(ctx, bson.D{{"_id", userID}}, bson.D{
//		{"$push", bson.D{{"refVideo", videoId}}},
//	})
//	if err != nil {
//		return err
//	}
//	return nil
//}
