package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func main() {
	_, _, collection := Connect()
	ctx := context.TODO()

	//how to work by go data structure in mongoDB

	var user1 User = User{
		Firstname: "hamid",
		Lastname:  "dehghanpour",
		Tags:      []string{"a", "b"},
		Age:       15,
	}
	var user2 User = User{
		Firstname: "ali",
		Lastname:  "dehghanpour",
		Tags:      []string{"a", "b"},
		Age:       20,
	}
	//create operation

	//create one item
	user, err := collection.InsertOne(ctx, user1)
	if err != nil {
		fmt.Println(err)
	}
	_ = user

	//create many item
	var usersInput []interface{}
	usersInput = append(usersInput, user1)
	usersInput = append(usersInput, user2)
	users, _ := collection.InsertMany(ctx, usersInput)
	_ = users
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Firstname string             `bson:"firstname,omitempty"`
	Lastname  string             `bson:"lastname,omitempty"`
	Tags      []string           `bson:"tags,omitempty"`
	Age       int                `bson:"age,omitempty"`
	Ref       primitive.ObjectID `bson:"ref,omitempty"`
}

func Connect() (*mongo.Client, *mongo.Database, *mongo.Collection) {
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
	//defer func() {
	//	if err = client.Disconnect(ctx); err != nil {
	//		panic(err)
	//	}
	//}()
	//send ping to mongo server for check connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("ping error")
	}
	//get certain database
	database := client.Database("testDatabase")
	//get certain collection in this database
	collection := database.Collection("user")
	_ = collection
	return client, database, collection
}
