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
	key := map[string]*dynamodb.AttributeValue{
		"id": &dynamodb.AttributeValue{
			S: aws.String("1"),
		},
	}
	//updateExpression :=expression.Set(expression.Name("name.id"),expression.Value([]string{"mohammad","ali"}))
	//updateExpression :=expression.Set(expression.Name("name.id[2]"),expression.Value("ali"))
	//updateExpression :=expression.Set(expression.Name("name"),expression.Value(map[string]string{"id":"1","name":"mohammad"}))
	//updateExpression := expression.Set(expression.Name("date"), expression.IfNotExists(expression.Name("date"), expression.Value("mohammad")))
	//updateExpression := expression.Set(expression.Name("name.id"),expression.ListAppend(expression.Name("name.id"),expression.Value([]string{"erfan"})))
	updateExpression := expression.Remove(expression.Name("date"))
	exp, err := expression.NewBuilder().WithUpdate(updateExpression).Build()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(exp.Names())
	fmt.Println(exp.Values())
	updateItemInput := &dynamodb.UpdateItemInput{
		TableName:                 aws.String("User"),
		Key:                       key,
		ReturnValues:              aws.String("ALL_NEW"),
		UpdateExpression:          exp.Update(),
		ExpressionAttributeNames:  exp.Names(),
		ExpressionAttributeValues: exp.Values(),
	}
	item, err := dynamoClient.UpdateItem(updateItemInput)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(item)
}
