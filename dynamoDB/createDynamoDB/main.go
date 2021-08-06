package createDynamoDB

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
)

func Dynamo()*dynamodb.DynamoDB  {
	err := os.Setenv("REGION", "us-west-2")
	if err != nil {
		return nil
	}
	region := os.Getenv("REGION")
	credential:= credentials.NewStaticCredentials("6heses","5mxn5","token")
	sessionConfig := aws.NewConfig().WithRegion(region).WithEndpoint("http://localhost:8000").WithCredentials(credential)
	awsSession, err := session.NewSession(sessionConfig)
	if err != nil {
		fmt.Println("error in create session")
		return nil
	}
	return dynamodb.New(awsSession)
}
