package main

import (
	"dynamoDB/createDynamoDB"
	"dynamoDB/data"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)
const TableName = "User"
func main()  {
	dynamoClient := createDynamoDB.Dynamo()
	user := data.User{
		Id: "2",
		Name: "a",
		Age: 22,
	}
	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(TableName),
		ReturnConsumedCapacity: aws.String("TOTAL"),
		ReturnValues: aws.String("ALL_OLD"),
	}
	a, err := dynamoClient.PutItem(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)

}


