package main

import (
	"dynamoDB/createDynamoDB"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

func main() {
	dynamoClient := createDynamoDB.Dynamo()

	cond:=expression.And(expression.Name("age").LessThanEqual(expression.Value(25)),expression.Name("id").Equal(expression.Value("1")))
	exp, err:=expression.NewBuilder().WithCondition(cond).Build()
	if err != nil {
		fmt.Println(err)
		return
	}
	deleteItemInput := &dynamodb.DeleteItemInput{
		TableName: aws.String("User"),
		Key: map[string]*dynamodb.AttributeValue{
			"id":&dynamodb.AttributeValue{
				S: aws.String("1"),
			},
		},
		ConditionExpression: exp.Condition(),
		ExpressionAttributeNames: exp.Names(),
		ExpressionAttributeValues: exp.Values(),
		ReturnValues: aws.String("ALL_OLD"),

	}
	item, err := dynamoClient.DeleteItem(deleteItemInput)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(item)
}
