package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/KingBean4903/graphql-vod-platform/auth"
	"github.com/KingBean4903/graphql-vod-platform/db"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/KingBean4903/graphql-vod-platform/graph"
	"github.com/vektah/gqlparser/v2/ast"


	"github.com/KingBean4903/graphql-vod-platform/internal/realtime"
)

const defaultPort = "8800"

func main() {
	port := os.Getenv("PORT")

	err := godotenv.Load()

	realtime := realtime.NewRedisPubSub()

	if port == "" {
		port = defaultPort
	}

	// Connect to Db
	db.Init()

	// GraphQL handler
	srv := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			PubSub: realtime,
		},
	}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	// Setup router
	r := mux.NewRouter()
	r.Use(auth.Middleware)


	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
