package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

//first step to create project by mongoDb
//mkdir goProject
//cd goProject
//go mod init goProject
//go get go.mongodb.org/mongo-driver/mongo

func main() {
	//define a timeout duration that we want to use when trying to connect.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//connect to mongoDB server
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	//check for error when connect to server
	if err != nil {
		fmt.Println("error")
	}
	//disconnect to mongo server when finish code execute
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	//send ping to mongo server for check connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("ping error")
	}
	//show databases name in mongo server
	databaseName, _ := client.ListDatabaseNames(ctx, bson.M{}) // show dbs
	_ = databaseName
	//get certain database
	database := client.Database("test")
	//get certain collection in this database
	collection := database.Collection("user")
	_ = collection

	//creat collection by schema validator
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"name", "age"},
		"properties": bson.M{
			"name": bson.M{
				"bsonType":    "string",
				"description": "the name of the user"},
			"age": bson.M{
				"bsonType":    "int",
				"minimum":     18,
				"description": "the age of the user"},
		},
	}
	validator := bson.M{"$jsonSchema": jsonSchema}
	opts := options.CreateCollection().SetValidator(validator)
	err = database.CreateCollection(ctx, "testUser", opts)
	if err != nil {
		fmt.Println("error")
	}
}
