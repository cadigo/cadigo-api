package app

import (
	"cadigo-api/config"
	"cadigo-api/db/mongodb/infrastructure"
	"cadigo-api/graph"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/caarlos0/env/v8"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

const defaultPort = "8080"

var mongodbConnector infrastructure.MongodbConnector
var generalConfig config.Config

func graphqlHandler() http.HandlerFunc {
	caddyHandler := caddyHandlerInit()
	bookinghandler := bookingHandlerInit()
	coursegolfhandler := courseGolfHandlerInit()
	customerhandler := customerHandlerInit()
	paymenthandler := paymentHandlerInit()
	chatHandler := chatHandlerInit()

	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CaddyHandler:      caddyHandler,
		BookingHandler:    bookinghandler,
		CoursegolfHandler: coursegolfhandler,
		CustomerHandler:   customerhandler,
		PaymentHandler:    paymenthandler,
		ChatHandler:       chatHandler,
	}}))

	h.AddTransport(&transport.Websocket{})

	return h.ServeHTTP
}

func playgroundHandler(endpoint string) http.HandlerFunc {
	h := playground.Handler("GraphQL", endpoint)

	return h.ServeHTTP
}

func loadConfig() {
	godotenv.Load()

	err := env.Parse(&generalConfig)
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}
}

func NewApp() error {
	loadConfig()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"Origin"},
	})

	r := mux.NewRouter()

	r.Handle("/", playgroundHandler("/graphql"))
	r.Handle("/graphql", c.Handler(graphqlHandler()))

	http.Handle("/", r)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "8080")

	log.Fatal(http.ListenAndServe(":8080", nil))

	return nil
}
