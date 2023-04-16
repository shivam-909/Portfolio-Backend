package response

import "github.com/aws/aws-lambda-go/events"

func New(body string, code int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Body:       string(body),
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
	}
}
