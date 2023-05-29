package app

import (
	"cadigo-api/app/handlers/paymenthandler"
	"cadigo-api/config"
	"cadigo-api/db/mongodb/infrastructure"
	"cadigo-api/graph"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/caarlos0/env/v8"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

const defaultPort = "8080"

var mongodbConnector infrastructure.MongodbConnector
var generalConfig config.Config
var paymentHandler *paymenthandler.Handler

func graphqlHandler() *handler.Server {
	caddyHandler := caddyHandlerInit()
	bookingHandler := bookingHandlerInit()
	coursegolfHandler := courseGolfHandlerInit()
	customerhandler := customerHandlerInit()
	paymentHandler = paymentHandlerInit()
	chatHandler := chatHandlerInit()

	h := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CaddyHandler:      caddyHandler,
		BookingHandler:    bookingHandler,
		CoursegolfHandler: coursegolfHandler,
		CustomerHandler:   customerhandler,
		PaymentHandler:    paymentHandler,
		ChatHandler:       chatHandler,
	}}))

	return h
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

func webSocketInit(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
	// Get the token from payload
	// any := initPayload["authToken"]
	// token, ok := any.(string)
	// if !ok || token == "" {
	// 	return nil, errors.New("authToken not found in transport payload")
	// }

	// // Perform token verification and authentication...
	userId := "john.doe" // e.g. userId, err := GetUserFromAuthentication(token)

	// put it in context
	ctxNew := context.WithValue(ctx, "username", userId)

	return ctxNew, nil
}

func NewApp() error {
	loadConfig()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},
		ExposedHeaders:   []string{"Origin"},
		AllowedHeaders:   []string{"Origin", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "X-Requested-With"},
		Debug:            false,
	})

	r := mux.NewRouter()

	r.Handle("/", playgroundHandler("/graphql"))

	// http.Handle("/", playgroundHandler("/graphql"))

	srv := graphqlHandler()
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	srv.Use(extension.Introspection{})
	r.Handle("/graphql", c.Handler(srv))

	// http.Handle("/graphql", c.Handler(srv))
	// http.Handle("/payment-confirm", c.Handler(srv))
	http.Handle("/", r)

	logrus.Info("connect to http://localhost:%s/ for GraphQL playground", "8080")

	logrus.Fatal(http.ListenAndServe(":8080", nil))

	return nil
}
