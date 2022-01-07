package gql

import (
	"demotoftheday.com/postgres"
	"github.com/graphql-go/graphql"
)

type Resolver struct {
	db *postgres.Db
}

func (r *Resolver) QuotesResolver(p graphql.ResolveParams) (interface{}, error) {
	quote := r.db.GetQuotes()

	return quote, nil
}

func (r *Resolver) RandomQuoteResolver(p graphql.ResolveParams) (interface{}, error) {
	quotes := r.db.GetRandomQuote()

	return quotes, nil
}
