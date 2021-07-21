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

//first step to create project by mongoDb
//mkdir goProject
//cd goProject
//go mod init goProject
//go get go.mongodb.org/mongo-driver/mongo
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Firstname string             `bson:"firstname,omitempty"`
	Lastname  string             `bson:"lastname,omitempty"`
	Tags      []string           `bson:"tags,omitempty"`
	Age       int                `bson:"age,omitempty"`
}

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
	collection := database.Collection("newUser")
	_ = collection
	// filtering
	filter := bson.D{{"$match", bson.D{{"firstname", "ali"}}}}
	pipeline := mongo.Pipeline{filter}
	user, _ := collection.Aggregate(ctx, pipeline)
	// group and sort
	group := bson.D{{"$group", bson.D{
		{"_id", "$firstname"},
		{"total", bson.D{{"$sum", 1}}},
	}}}
	sort := bson.D{{"$sort", bson.D{{"total", -1}}}}

	pipeline = mongo.Pipeline{group, sort}

	user, _ = collection.Aggregate(ctx, pipeline)

	// project
	project := bson.D{{
		"$project",
		bson.D{
			{"_id", 0}, // don't show _id

			{"firstname", 1}, // show firstname

			{ //convert type to other type
				"age",
				bson.D{{
					"$convert", // or use $toInt operator
					bson.D{
						{"input", "$age"},
						{"to", "int"},
					},
				}},
			}, //convert type to other type

			{"fullName", bson.D{{"$concat", bson.A{ // concat multiple item
				bson.D{ // convert to upper case
					{"$toUpper", bson.D{{"$substrCP", bson.A{"$firstname", 0, 1}}}},
				},
				bson.D{{ // retrieve section of strign
					"$substrCP",
					bson.A{
						"$firstname",
						1,
						bson.D{{ // subtract operation
							"$subtract",
							bson.A{
								bson.D{{"$strLenCP", "$firstname"}}, 1},
						}},
					},
				}},
				" ",
				"$lastname"}}}},
		},
	}}
	skip := bson.D{{"$skip", 1}}
	limit := bson.D{{"$limit", 1}}
	pipeline = mongo.Pipeline{project, skip, limit}

	lookup := bson.D{{ // for lookup reference
		"$lookup",
		bson.D{
			{"from", "abc"},
			{"localField", "ref"},
			{"foreignField", "_id"},
			{"as", "show"},
		},
	}}
	pipeline = mongo.Pipeline{lookup}

	user, _ = collection.Aggregate(ctx, pipeline)
	fmt.Println(user)
	var userDecode []interface{}
	err = user.All(ctx, &userDecode)
	if err != nil {
		return
	}
	fmt.Println(userDecode)
}
