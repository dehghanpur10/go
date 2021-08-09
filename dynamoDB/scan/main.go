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
	scanInput := &dynamodb.ScanInput{
		ExpressionAttributeValues: expr.Values(),
		ExpressionAttributeNames: expr.Names(),
		FilterExpression: expr.Filter(),
		IndexName: aws.String("age-index"),
		TableName: aws.String("User"),
	}
	scan, err := dynamoClient.Scan(scanInput)
	if err != nil {
		return
	}
	fmt.Println(scan)
}
