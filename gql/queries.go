package gql

import (
	"demotoftheday.com/postgres"
	"github.com/graphql-go/graphql"
)

type Root struct {
	Query *graphql.Object
}

func NewRoot(db *postgres.Db) *Root {
	resolver := Resolver{db: db}

	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"quotes": &graphql.Field{
						Type:    graphql.NewList(Quote),
						Resolve: resolver.QuotesResolver,
					},
					"randomQuote": &graphql.Field{
						Type:    Quote,
						Resolve: resolver.RandomQuoteResolver,
					},
				},
			},
		),
	}

	return &root
}
