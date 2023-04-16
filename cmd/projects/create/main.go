package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/shivam-909/portfolio-backend/internal/dynamo"
	"github.com/shivam-909/portfolio-backend/pkg/projects"
	"github.com/shivam-909/portfolio-backend/pkg/response"

	dynamo_projects "github.com/shivam-909/portfolio-backend/internal/dynamo/projects"
)

func main() {
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	sess := dynamo.CreateDynamoSession()

	project, err := projects.Parse(req)
	if err != nil {
		return response.New(err.Error(), http.StatusBadRequest), err
	}

	err = dynamo_projects.CreateProject(sess, project)
	if err != nil {
		return response.New(err.Error(), http.StatusInternalServerError), err
	}

	return response.New("", http.StatusAccepted), nil
}
