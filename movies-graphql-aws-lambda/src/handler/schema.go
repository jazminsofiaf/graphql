package handler
var Schema = `
	schema {
		query: Query
	}
	type Movie {
		title: String!
		plot: String
		rating: Float
		year: Int
		
	}
	type Query{
		movie(title: String! rating:Float): Movie
	}
`
