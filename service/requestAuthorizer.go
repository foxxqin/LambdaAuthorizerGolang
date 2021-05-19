package service

import (
	"context"
	"errors"
	"fmt"
	"lambdaAPIAuthorizer/database"
	"lambdaAPIAuthorizer/model"
	"lambdaAPIAuthorizer/utility"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
)

func RequestAuthorizer(request events.APIGatewayCustomAuthorizerRequestTypeRequest) (bool, error) {
	headers := request.Headers
	queryString := request.QueryStringParameters

	// log.Printf("ClientID: %s", headers["client_id"])
	// log.Printf("Signature: %s", headers["signature"])
	// log.Printf("RequestID: %s", queryString["request_id"])
	// log.Printf(request.MethodArn)

	if headers["client_id"] == "" || headers["signature"] == "" || queryString["request_id"] == "" {
		err := errors.New("Missing Authentication Information")
		return false, err
	} else {

		clientSecret, err := getClientKeyByClientID(headers["client_id"])
		if err != nil || clientSecret == "" {
			return false, err
		}

		plain := fmt.Sprintf("%s=%s&%s=%s&%s=%s", "client_id", headers["client_id"],
			"client_key", clientSecret, "request_id", queryString["request_id"])

		// log.Printf("Plain: %s", plain)
		hash := utility.GetMD5Hash(plain)
		// log.Printf("Hash: %s", hash)

		if hash != headers["signature"] {
			err := errors.New("Ivalid Signature")
			return false, err
		}
	}
	return true, nil
}

func getClientKeyByClientID(clientID string) (string, error) {
	db, err := database.GetDynamoDBConnection()

	if err != nil {
		log.Printf("getdatabase error: %s", err.Error())
		return "", err
	}

	var searchKey model.SearchKey
	searchKey.ClientID = clientID
	key, err := attributevalue.MarshalMap(searchKey)

	if err != nil {
		log.Printf("marshalmap error: %s", err.Error())
		return "", err
	}

	input := &dynamodb.GetItemInput{
		TableName:            aws.String("ApiUser"),
		ProjectionExpression: aws.String("ClientSecret"),
		Key:                  key,
	}

	result, err := db.GetItem(context.TODO(), input)

	if err != nil || result == nil {
		log.Printf("get item error: %s", err.Error())
		return "", err
	}

	var clientSecretItem model.ClientSecretItem
	err = attributevalue.UnmarshalMap(result.Item, &clientSecretItem)

	if err != nil {
		log.Printf("unmarshal map error: %s", err.Error())
		return "", err
	}

	return clientSecretItem.ClientSecret, nil
}
