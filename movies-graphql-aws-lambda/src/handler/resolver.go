package handler

import (
	"github.com/graph-gophers/graphql-go"
	"jaz.com/movies-graphql/movies-graphql-aws-lambda/src/domain"
	"log"
)


type Resolver struct{}

var moviesData = make(map[string]*domain.Item)

var movies = []* domain.Item{
	{
		Year : 2013,
		Title : "Turn It Down, Or Else!",
		Plot : "A rock band plays their music at high volumes, annoying the neighbors.",
		Rating : 4.5,
	},
	{
		Year: 2015,
		Title: "The Big New Movie",
		Plot: "Nothing happens at all.",
		Rating: 0.1,
	},
	{
		Year: 1994,
		Title: "La sirenita",
		Plot: "Una joven sirena se enamora de un humano",
		Rating: 5.0,
	},
}

var mainSchema *graphql.Schema

func init() {
	for _, p := range movies {
		moviesData[p.Title] = p
	}
}

func (r *Resolver) Movie(args struct{ Title string }) *domain.ItemResolver {
	if i := moviesData[args.Title]; i != nil {
		log.Print("Found in resolver!/n")
		return &domain.ItemResolver{I: i}
	}
	return nil
}
