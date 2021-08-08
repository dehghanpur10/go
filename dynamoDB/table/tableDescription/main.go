package main

import (
	"dynamoDB/createDynamoDB"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	dynamoClient := createDynamoDB.Dynamo()
	describeTable := &dynamodb.DescribeTableInput{
		TableName: aws.String("User"),
	}
	table, err := dynamoClient.DescribeTable(describeTable)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(table)
}
