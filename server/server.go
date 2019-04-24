package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	hello_world "github.com/graphql-services/hello-world"
)

const defaultPort = "80"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/graphql"))
	http.Handle("/graphql", handler.GraphQL(hello_world.NewExecutableSchema(hello_world.Config{Resolvers: &hello_world.Resolver{}})))

	http.HandleFunc("/healthcheck", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("OK"))
		res.WriteHeader(200)
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
