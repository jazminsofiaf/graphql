package main

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"jaz.com/movies-graphql/movies-graphql-aws-lambda/src/handler"
	"log"
	"net/http"
)

var Schema = `
	schema {
		query: Query
	}
	type Person {
		firstName: String!
		lastName: String
	}
	type Query{
		person(id: ID!): Person
	}
`

var mainSchema *graphql.Schema

func init() {
	mainSchema = graphql.MustParseSchema(Schema, &handler.Resolver{})
}

func main() {
	http.Handle("/query", &relay.Handler{Schema: mainSchema})
	log.Print("Starting to listen 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

	//input :=  "The Big New Movie"
	//handler.HandleRequest(nil, input)
	//lambda.Start(handler.HandleRequest)
}