package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"demotoftheday.com/gql"
	"demotoftheday.com/postgres"
	"demotoftheday.com/server"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
)

func main() {
	router, db := initializeApp()
	defer db.Close()

	log.Fatal(http.ListenAndServe(":3000", router))
}

func initializeApp() (*chi.Mux, *postgres.Db) {
	router := chi.NewRouter()

	db, err := postgres.New(postgres.ConnString(
		os.Getenv("POSTGRE_HOST"),
		os.Getenv("POSTGRE_PORT"),
		os.Getenv("POSTGRE_USER"),
		os.Getenv("POSTGRE_PASSWORD"),
		os.Getenv("POSTGRE_DB_NAME")))

	if err != nil {
		log.Fatal(err)
	}

	rootQuery := gql.NewRoot(db)

	sc, err := graphql.NewSchema(graphql.SchemaConfig{Query: rootQuery.Query})
	if err != nil {
		fmt.Println("Error creating schema:", err)
	}

	s := server.Server{GqlSchema: &sc}

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.Compress(5),
		middleware.StripSlashes,
		middleware.Recoverer,
	)

	router.Post("/graphql", s.GraphQL())
	router.Get("/", s.Home(db))

	return router, db
}
