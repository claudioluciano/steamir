package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/claudioluciano/steamir/api/graph"
	"github.com/claudioluciano/steamir/api/graph/generated"
	"github.com/claudioluciano/steamir/api/internal/pkg/steamclient"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	steamKey := os.Getenv("STEAM_KEY")

	steamclient := steamclient.New(steamKey)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{
			Resolvers: &graph.Resolver{
				SteamClient: steamclient,
			},
		}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
