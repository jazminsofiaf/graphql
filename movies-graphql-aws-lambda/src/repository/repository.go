package repository

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"jaz.com/movies-graphql/movies-graphql-aws-lambda/src/domain"
)

const (
	tableName = "Movies"
)

type MoviesRepository struct {
	TableName string
	DynamoClient *dynamodb.DynamoDB
}

func NewMoviesRepository(client *dynamodb.DynamoDB) *MoviesRepository {
	repository := new(MoviesRepository)
	repository.TableName = tableName
	repository.DynamoClient = client
	return repository
}

func (repository MoviesRepository) GetMovie(movieName string, movieYear string) (domain.Item, error) {
	movie := domain.Item{}

	result, err := repository.DynamoClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(repository.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Year": {
				N: aws.String(movieYear),
			},
			"Title": {
				S: aws.String(movieName),
			},
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return movie, err
	}

	if result.Item == nil {
		msg := "Could not find '" + movieName + "'"
		return movie, errors.New(msg)
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &movie)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	return movie, nil
}


