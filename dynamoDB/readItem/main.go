package main

import (
	"dynamoDB/createDynamoDB"
	"dynamoDB/data"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"

)

func main() {
	dynamoClient := createDynamoDB.Dynamo()

	Key := map[string]*dynamodb.AttributeValue{
		"id": {
			S: aws.String("2"),
		},
	}
	proj := expression.NamesList(expression.Name("age"),expression.Name("name"))
	expr ,err := expression.NewBuilder().WithProjection(proj).Build()
	if err != nil {
		fmt.Println(err)
	}
	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String("User"),
		Key:       Key,
		ExpressionAttributeNames: expr.Names(),
		ProjectionExpression: expr.Projection(),
		//ProjectionExpression: aws.String("age,#N"),
		//ExpressionAttributeNames: map[string]*string{
		//	"#N":aws.String("name"),
		//},
	}
	result,err:= dynamoClient.GetItem(getItemInput)
	if err != nil {
		fmt.Println(err)
		return
	}
	if result.Item == nil {
		fmt.Println("user not found")
		return
	}
	var user data.User
	err = dynamodbattribute.UnmarshalMap(result.Item, &user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}
