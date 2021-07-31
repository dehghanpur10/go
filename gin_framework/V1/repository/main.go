package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database DatabaseHelper

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

func NewDatabase(url string, databaseName string) error {
	c, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		return err
	}
	client := mongoClient{client: c}
	Database = client.Database(databaseName)
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
