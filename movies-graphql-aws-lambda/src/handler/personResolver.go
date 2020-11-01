package handler

type Person struct {
	ID        string
	FirstName string
	LastName  string
}

type PersonResolver struct {
	P *Person
}

func (r *PersonResolver) ID() string {
	return r.P.ID
}

func (r *PersonResolver) FirstName() string {
	return r.P.FirstName
}

func (r *PersonResolver) LastName() *string {
	return &r.P.LastName
}
