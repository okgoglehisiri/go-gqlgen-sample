import (
	"github.com/okgoglehisiri/graphql-sample/graph"
	"github.com/okgoglehisiri/graphql-sample/internal"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(internal.NewExecutableSchema(internal.config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	Http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s for GraphQL playground", port)
	log.Printf(http.ListenAndServe(":"+port, nil))
}