package gql

import "github.com/graphql-go/graphql"

var Quote = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Quote",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"quotation": &graphql.Field{
				Type: graphql.String,
			},
			"person": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
