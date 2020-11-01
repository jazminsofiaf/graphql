package repository

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"jaz.com/movies-graphql/movies-graphql-aws-lambda/src/domain"
	"os"
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

func (repository MoviesRepository)  GetAllMoviesWithName(title string) ([]domain.Item, error) {
	// Create the Expression to fill the input struct with.
	// Get all movies in that year; we'll pull out those with a higher rating later
	filter := expression.Name("Title").Equal(expression.Value(title))

	// Get back the title, year, and rating
	proj := expression.NamesList(expression.Name("Title"))
	proj = proj.AddNames(expression.Name("Year"))
	proj = proj.AddNames(expression.Name("Plot"))
	proj = proj.AddNames(expression.Name("Rating"))

	expr, err := expression.NewBuilder().WithFilter(filter).WithProjection(proj).Build()
	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(repository.TableName),
	}


	fmt.Println("Make the DynamoDB Query API call")
	// Make the DynamoDB Query API call
	result, err := repository.DynamoClient.Scan(params)
	if err != nil {
		fmt.Println("Query API call failed:")
		fmt.Println((err.Error()))
		os.Exit(1)
	}


	moviesSlice := make([]domain.Item, 0)
	numItems := 0
	for _, itemStruct := range result.Items {
		item := domain.Item{}

		err = dynamodbattribute.UnmarshalMap(itemStruct, &item)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			os.Exit(1)
		}


		fmt.Println()
		fmt.Println("Title: ", item.Title)
		fmt.Println("Rating:", item.Rating)
		fmt.Println("Plot:", item.Plot)
		moviesSlice = append(moviesSlice, item)

	}

	fmt.Println("Found", numItems, "movie(s) with a title", title)
	return moviesSlice, nil
}


