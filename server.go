package main

import (
	"log"
	"my-go-apps/VitaminApp/database"
	"my-go-apps/VitaminApp/graph"
	"my-go-apps/VitaminApp/graph/generated"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"my-go-apps/VitaminApp/cors"
)

const defaultPort = "43341"

func main() {
	database.SetupDatabase()
	// go client()
	// go server()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", cors.MiddlewareHandler(srv))


	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Println(http.ListenAndServe(":"+port, nil))

}
