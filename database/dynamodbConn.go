package database

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func GetDynamoDBConnection() (*dynamodb.Client, error) {
	config, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		return nil, err
	}

	client := dynamodb.NewFromConfig(config)
	return client, nil

}
