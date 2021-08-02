package videoCollection

//
//import (
//	"V1/models"
//	"V1/repository"
//	"context"
//	"fmt"
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	"go.mongodb.org/mongo-driver/mongo"
//)
//
//const collectionName = "videos"
//
//type UserDatabase interface {
//	FindOne(ctx context.Context, id string) (*models.Video, error)
//	Create(ctx context.Context, usr *models.VideoByAuthorId) (interface{}, error)
//}
//
//type userDatabase struct {
//	db repository.DatabaseHelper
//}
//
//func NewVideoDatabase(db repository.DatabaseHelper) UserDatabase {
//	return &userDatabase{
//		db: db,
//	}
//}
//
//func (u *userDatabase) FindOne(ctx context.Context, id string) (*models.Video, error) {
//	objectId, err := primitive.ObjectIDFromHex(id)
//	if err != nil {
//		return nil, err
//	}
//	filter := bson.D{{"$match", bson.D{{"_id", objectId}}}}
//	lookup := bson.D{{ // for lookup reference
//		"$lookup",
//		bson.D{
//			{"from", "users"},
//			{"localField", "author"},
//			{"foreignField", "_id"},
//			{"as", "author"},
//		},
//	}}
//	pipeline := mongo.Pipeline{filter, lookup}
//	aggregate, err := u.db.Collection(collectionName).Aggregate(context.TODO(), pipeline)
//	if err != nil {
//		return nil, err
//	}
//	var video []models.Video
//	err = aggregate.All(ctx, &video)
//	if err != nil {
//		fmt.Println(err)
//		fmt.Println("fdfd")
//		return nil, err
//	}
//	fmt.Println(video)
//	return nil, nil
//}
//
//func (u *userDatabase) Create(ctx context.Context, video *models.VideoByAuthorId) (interface{}, error) {
//	newUser, err := u.db.Collection(collectionName).InsertOne(ctx, video)
//	if err != nil {
//		return nil, err
//	}
//	return newUser.InsertedID, nil
//}
//
