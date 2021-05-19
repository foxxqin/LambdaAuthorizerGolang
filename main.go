package main

import (
	"context"
	"lambdaAPIAuthorizer/model"
	"lambdaAPIAuthorizer/service"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequestSimple(ctc context.Context, request events.APIGatewayCustomAuthorizerRequestTypeRequest) (model.Response, error) {

	var response model.Response

	verifyResult, err := service.RequestAuthorizer(request)

	if verifyResult == true {
		response.IsAuthorized = true
	}

	return response, err
}

func main() {
	lambda.Start(HandleRequestSimple)
}
