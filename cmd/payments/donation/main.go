package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/shivam-909/portfolio-backend/pkg/response"
)

func main() {
	lambda.Start(handler)
}

type Request struct{}

// TODO(shivam): Implement
func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var request Request
	err := json.Unmarshal([]byte(req.Body), &request)
	if err != nil {
		return response.New(err.Error(), http.StatusBadRequest), err
	}

	return response.New("", http.StatusAccepted), nil
}
