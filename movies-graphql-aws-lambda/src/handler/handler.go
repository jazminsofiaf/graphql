package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"jaz.com/movies-graphql/movies-graphql-aws-lambda/src/client"
	"jaz.com/movies-graphql/movies-graphql-aws-lambda/src/domain"
	"jaz.com/movies-graphql/movies-graphql-aws-lambda/src/repository"
	"log"
)

var moviesRepository *repository.MoviesRepository
//var mainSchema *graphql.Schema

func init(){
	client, err := client.NewLocalClient()
	if err != nil {
		errors.New("error at creating Client")
	}
	moviesRepository = repository.NewMoviesRepository(client)
}


func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, errors.New("no query was provided in the HTTP body")
	}

	if err := json.Unmarshal([]byte(request.Body), &requestParams); err != nil {
		log.Print("Could not decode body", err)
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       "",
		StatusCode: 200,
	}, nil
}

type Handler struct{
}



func (h Handler) Movie(args struct{ Title string; Rating *float64}) *domain.ItemResolver {
	fmt.Println("------------------- all movies --------------")
	movies, _:= moviesRepository.GetAllMoviesWithName(args.Title)

	if args.Rating == nil {
		if i := &movies[0]; i != nil {
			log.Println("First movie with that title!")
			return &domain.ItemResolver{I: i}
		}
	}

	moviesByRatingMap := make(map[float64]*domain.Item)
	for _, m := range movies {
		moviesByRatingMap[m.Rating] = &m
	}
	if i := moviesByRatingMap[*args.Rating]; i != nil {
		log.Println("Found in resolver!")
		return &domain.ItemResolver{I: i}
	}
	return nil
}
