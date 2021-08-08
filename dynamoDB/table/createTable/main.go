package main

import (
	"dynamoDB/createDynamoDB"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	dynaClient := createDynamoDB.Dynamo()
	tableInput := createTableInput()
	table, err := dynaClient.CreateTable(tableInput)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(table)
}
func createTableInput() *dynamodb.CreateTableInput {
	// Create table Movies
	tableName := "User"

	return &dynamodb.CreateTableInput{
		LocalSecondaryIndexes: localSecondaryIndex(),
		AttributeDefinitions:  attributeDefinitions(),
		KeySchema:             keySchemaBySortKey("name"),
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}
}
func attributeDefinitions() []*dynamodb.AttributeDefinition {
	return []*dynamodb.AttributeDefinition{
		{
			AttributeName: aws.String("id"),
			AttributeType: aws.String("S"),
		},
		{
			AttributeName: aws.String("name"),
			AttributeType: aws.String("S"),
		},
		{
			AttributeName: aws.String("age"),
			AttributeType: aws.String("N"),
		},
	}
}
func keySchemaBySortKey(sortKey string) []*dynamodb.KeySchemaElement {
	return []*dynamodb.KeySchemaElement{
		{
			AttributeName: aws.String("id"),
			KeyType:       aws.String("HASH"),
		},
		{
			AttributeName: aws.String(sortKey),
			KeyType:       aws.String("RANGE"),
		},
	}
}
func localSecondaryIndex() []*dynamodb.LocalSecondaryIndex {
	return []*dynamodb.LocalSecondaryIndex{
		{
			IndexName: aws.String("age-index"),
			KeySchema: keySchemaBySortKey("age"),
			Projection: &dynamodb.Projection{
				ProjectionType: aws.String("ALL"),
			},
		},
	}
}
