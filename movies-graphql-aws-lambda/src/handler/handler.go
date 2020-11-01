package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"jaz.com/movies-graphql/movies-graphql-aws-lambda/src/client"
	"jaz.com/movies-graphql/movies-graphql-aws-lambda/src/repository"
	"log"
)


func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, errors.New("no query was provided in the HTTP body")
	}

	if err := json.Unmarshal([]byte(request.Body), &requestParams); err != nil {
		log.Print("Could not decode body", err)
		return events.APIGatewayProxyResponse{}, err
	}

	client, err := client.NewLocalClient()
	if err != nil {
		return events.APIGatewayProxyResponse{}, errors.New("error at creating Client")
	}
	moviesRepository := repository.NewMoviesRepository(client)
	handler := NewHandler(moviesRepository)
	handler.Handle("The Big New Movie", "2016")


	return events.APIGatewayProxyResponse{
		Body:       "",
		StatusCode: 200,
	}, nil
}

type Handler struct{
	moviesRepository repository.MoviesRepository
}

func NewHandler(repository *repository.MoviesRepository) *Handler {
	this := new(Handler)
	this.moviesRepository = *repository
	return this
}


func (h Handler) Handle(movieName string , movieYear string)   {


	fmt.Print("------------------- looking for: ")
	fmt.Print(movieName)
	fmt.Println(movieYear)

	movie, err := h.moviesRepository.GetMovie(movieName, movieYear)
	if err != nil {
		errors.New(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	fmt.Println("------------------- Found movie return answer --------------")
	fmt.Println("Title: ", movie.Title)
	fmt.Println("Plot: ", movie.Plot)
	fmt.Println("Year: ", movie.Year)
	fmt.Println("Rating:", movie.Rating)

}
