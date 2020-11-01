package handler

import "github.com/graph-gophers/graphql-go"

type Person struct {
	ID        graphql.ID
	FirstName string
	LastName  string
}

type PersonResolver struct {
	P *Person
}

func (r *PersonResolver) ID() graphql.ID {
	return r.P.ID
}

func (r *PersonResolver) FirstName() string {
	return r.P.FirstName
}

func (r *PersonResolver) LastName() *string {
	return &r.P.LastName
}
