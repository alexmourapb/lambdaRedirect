package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusPermanentRedirect,
			Headers: map[string]string{
				"location": "https://www.google.com",
			},
		}, nil
}

func main() {
	lambda.Start(Handler)
}
