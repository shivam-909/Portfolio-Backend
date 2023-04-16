package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/shivam-909/portfolio-backend/internal/dynamo"
	"github.com/shivam-909/portfolio-backend/internal/projects"
)

func main() {
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	sess := createDynamoSession()

	var project projects.Project

	err := json.Unmarshal([]byte(req.Body), &project)
	if err != nil {
		return response(err.Error(), 400), err
	}

	err = dynamo.CreateProject(sess, project)
	if err != nil {
		return response(err.Error(), 500), err
	}

	return response("", 200), nil
}

func createDynamoSession() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(
		session.Options{
			SharedConfigState: session.SharedConfigEnable,
		},
	))

	return dynamodb.New(sess)
}

func response(body string, code int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Body:       string(body),
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
	}
}
