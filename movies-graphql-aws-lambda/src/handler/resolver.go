package handler

import (
	"github.com/graph-gophers/graphql-go"
	"log"
)


type Resolver struct{}

var peopleData = make(map[string]*Person)

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
		peopleData[p.FirstName] = p
	}
}

func (r *Resolver) Person(args struct{ FirstName string }) *PersonResolver {
	if p := peopleData[args.FirstName]; p != nil {
		log.Print("Found in resolver!/n")
		return &PersonResolver{p}
	}
	return nil
}
