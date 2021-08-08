package main

import (
	"dynamoDB/createDynamoDB"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main()  {
	dynamoClient := createDynamoDB.Dynamo()
	scanInput := &dynamodb.QueryInput{

	}
	scan, err := dynamoClient.Query(scanInput)
	if err != nil {
		return
	}
	_=scan
}
