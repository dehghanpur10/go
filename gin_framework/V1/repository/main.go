package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client ClientHelper

type DatabaseHelper interface {
	Collection(name string) *mongo.Collection
}
type ClientHelper interface {
	Database(string) DatabaseHelper
}

type mongoClient struct {
	client *mongo.Client
}
type mongoDatabase struct {
	db *mongo.Database
}

func NewClient(url string) error {
	c, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	Client = &mongoClient{client: c}
	return nil
}

func (mc *mongoClient) Database(dbName string) DatabaseHelper {
	db := mc.client.Database(dbName)
	return &mongoDatabase{db: db}
}

func (md *mongoDatabase) Collection(colName string) *mongo.Collection {
	collection := md.db.Collection(colName)
	return collection
}
