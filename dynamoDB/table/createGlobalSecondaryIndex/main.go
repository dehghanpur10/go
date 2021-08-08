package main

import (
	"dynamoDB/createDynamoDB"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	dynamoClient := createDynamoDB.Dynamo()
	updateTableInput := &dynamodb.UpdateTableInput{
		TableName:                   aws.String("User"),
		AttributeDefinitions:        attributeDefinition("job", "S", "age", "N"),
		GlobalSecondaryIndexUpdates: globalSecondaryIndex("job-index", "job", "age"),
	}
	table, err := dynamoClient.UpdateTable(updateTableInput)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(table)
}
func attributeDefinition(partitionKey, partitionKeyType, sortKey, sortKeyType string) []*dynamodb.AttributeDefinition {
	return []*dynamodb.AttributeDefinition{
		{
			AttributeName: aws.String(partitionKey),
			AttributeType: aws.String(partitionKeyType),
		},
		{
			AttributeName: aws.String(sortKey),
			AttributeType: aws.String(sortKeyType),
		},
	}
}
func globalSecondaryIndex(indexName, partitionKEy, sortKey string) []*dynamodb.GlobalSecondaryIndexUpdate {
	return []*dynamodb.GlobalSecondaryIndexUpdate{
		{
			Create: &dynamodb.CreateGlobalSecondaryIndexAction{
				IndexName: aws.String(indexName),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String(partitionKEy),
						KeyType:       aws.String("HASH"),
					},
					{
						AttributeName: aws.String(sortKey),
						KeyType:       aws.String("RANGE"),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("ALL"),
				},
			},
		},
	}
}
