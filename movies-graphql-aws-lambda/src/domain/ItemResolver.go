package domain
type ItemResolver struct {
	I *Item
}

func (r *ItemResolver) Year() *int32 {
	y := int32(r.I.Year)
	return &y
}

func (r *ItemResolver) Title() string {
	return r.I.Title
}

func (r *ItemResolver) Plot() *string {
	return &r.I.Plot
}

func (r *ItemResolver) Rating() *float64 {
	return &r.I.Rating
}
