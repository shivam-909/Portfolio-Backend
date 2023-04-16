package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/shivam-909/portfolio-backend/internal/dynamo"
	dynamo_projects "github.com/shivam-909/portfolio-backend/internal/dynamo/projects"
	"github.com/shivam-909/portfolio-backend/pkg/response"
)

func main() {
	lambda.Start(handler)
}

type Request struct {
	ID int32 `json:"id"`
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	sess := dynamo.CreateDynamoSession()

	var request Request
	err := json.Unmarshal([]byte(req.Body), &request)
	if err != nil {
		return response.New(err.Error(), http.StatusBadRequest), err
	}

	project, err := dynamo_projects.RetrieveProject(sess, request.ID)
	if err != nil {
		return response.New(err.Error(), http.StatusInternalServerError), err
	}

	return response.New(project.String(), http.StatusAccepted), nil
}
