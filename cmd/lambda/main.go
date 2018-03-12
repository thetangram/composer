package main

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Application build details
var (
	version   = "development"
	build     = "undefined"
	buildDate = "undefined"
)

func main() {
	fmt.Println("The Tangram Composer Lambda launcher")
	log.Printf("\tversion:      %s\n", version)
	log.Printf("\tbuild:        %s\n", build)
	log.Printf("\tbuild date:   %s\n", buildDate)
	log.Printf("\tstartup date: %s\n", time.Now().Format(time.RFC3339))

	lambda.Start(Handler)
}

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambdevents.APIGatewayProxyResponsea request %s\n", request.RequestContext.RequestID)
	return events.APIGatewayProxyResponse{
		Body:       "Not implemented yet",
		StatusCode: 200,
	}, nil
}
