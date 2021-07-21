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

	index := []mongo.IndexModel{
		{
			Keys: bson.D{{"firstname", "text"}, {"lastname", "text"}},
			Options: options.Index().SetUnique(true).SetWeights(bson.D{
				{"firstname", 2},
				{"lastname", 5},
			}), // set unique
		},
		{
			Keys: bson.D{{"age", 1}},
			Options: options.Index().SetPartialFilterExpression(bson.D{{ // create partial filter index
				"age",
				bson.D{{"$gt", 20}},
			},
			}),
		},
		{
			Keys:    bson.D{{"firstname", 1}},
			Options: options.Index().SetExpireAfterSeconds(2),
		},
	}
	many, err := collection.Indexes().CreateMany(ctx, index)
	if err != nil {

	}
	fmt.Println(many)

	//find document by text index
	user, _ := collection.Find(ctx, bson.D{{"$text", bson.D{{"$search", "dehghanpour"}}}})
	var a []interface{}
	user.All(ctx, &a)
	fmt.Println(a)

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
