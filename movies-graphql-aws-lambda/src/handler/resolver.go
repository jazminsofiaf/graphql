package handler

import (
	"github.com/graph-gophers/graphql-go"
	"log"
)


type Resolver struct{}

var peopleData = make(map[graphql.ID]*Person)

var people = []* Person{
	{
		ID:        "1000",
		FirstName: "Pedro",
		LastName:  "Marquez",
	},

	{
		ID:        "1001",
		FirstName: "John",
		LastName:  "Doe",
	},
}

var mainSchema *graphql.Schema

func init() {
	for _, p := range people {
		peopleData[p.ID] = p
	}
}

func (r *Resolver) Person(args struct{ ID graphql.ID }) *PersonResolver {
	if p := peopleData[args.ID]; p != nil {
		log.Print("Found in resolver!/n")
		return &PersonResolver{p}
	}
	return nil
}
