package main

import (
	"dynamoDB/createDynamoDB"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

func main()  {
	dynamoClient := createDynamoDB.Dynamo()
	filter := expression.Name("age").LessThanEqual(expression.Value(22)).And(expression.Name("id").Equal(expression.Value("5")))
	expr,err := expression.NewBuilder().WithFilter(filter).Build()
	if err != nil {
		fmt.Println(err)
		return
	}
	queryInput := &dynamodb.QueryInput{
		ExpressionAttributeValues: expr.Values(),
		ExpressionAttributeNames: expr.Names(),
		KeyConditionExpression: expr.Filter(),
		IndexName: aws.String("age-index"),
		TableName: aws.String("User"),
	}
	query, err := dynamoClient.Query(queryInput)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(query)
}
