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
		Name: "mohammad",
		Age: 22,
	}
	av, err := dynamodbattribute.MarshalMap(user)
	_ =av
	if err != nil {
		fmt.Println(err)
		return
	}

	input := &dynamodb.PutItemInput{
		Item:    av,
		TableName: aws.String(TableName),
		ReturnConsumedCapacity: aws.String("TOTAL"),
		ReturnValues: aws.String("ALL_OLD"),
		ConditionExpression: aws.String("attribute_not_exists(id)"),
	}
	a, err := dynamoClient.PutItem(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(a)

}


