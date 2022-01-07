package server

import (
	"encoding/json"
	"html/template"
	"net/http"

	"demotoftheday.com/gql"
	"demotoftheday.com/postgres"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
)

type Server struct {
	GqlSchema *graphql.Schema
}

type reqBody struct {
	Query string `json:"query"`
}

func (s *Server) GraphQL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			http.Error(w, "Must provide graphql query in request body", 400)
			return
		}

		var rBody reqBody
		err := json.NewDecoder(r.Body).Decode(&rBody)
		if err != nil {
			http.Error(w, "Error parsingJSON request body", 400)
		}

		result := gql.ExecuteQuery(rBody.Query, *s.GqlSchema)

		render.JSON(w, r, result)
	}
}

func (s *Server) Home(db *postgres.Db) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("views/home.html")
		t.Execute(w, db.GetRandomQuote())
	}
}
