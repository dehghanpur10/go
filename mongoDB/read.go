package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func main() {
	_, _, collection := Connect()
	ctx := context.TODO()

	//read operation

	//find one item
	var fetchUser User
	fetch := collection.FindOne(ctx, bson.M{"firstname": "hamid"})
	if err := fetch.Decode(&fetchUser); err != nil {
		panic(err)
	}
	//find many item
	var fetchUsers []User
	o := options.Find()
	o.SetProjection(bson.D{{"_id", 0}, {"firstname", 1}})
	o.SetLimit(2)
	o.SetSkip(1)
	o.SetSort(bson.D{{"age", -1}})

	filter := bson.D{
		// aga > 10 && age < 21 && (firstname == hamid && firstname == mohammad)
		{"age", bson.D{{"$gt", 10}}},
		{"age", bson.D{{"$lt", 21}}},
		{"$or", bson.A{
			bson.D{{"firstname", "mohammad"}},
			bson.D{{"firstname", "hamid"}},
		}},
	}
	userFind, _ := collection.Find(ctx, filter, o)

	if err := userFind.All(ctx, &fetchUsers); err != nil {
	}

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
