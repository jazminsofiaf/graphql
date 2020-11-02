package main

import (
	"jaz.com/movies-graphql/movies-graphql-aws-lambda/src/handler"
)

func main() {
	handler.StartLocally()
	//lambda.Start(handler.HandleRequest)
}