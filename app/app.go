package app

import (
	"cadigo-api/config"
	"cadigo-api/db/mongodb/infrastructure"
	"cadigo-api/graph"
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/caarlos0/env/v8"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

var mongodbConnector infrastructure.MongodbConnector
var generalConfig config.Config

func graphqlHandler() gin.HandlerFunc {
	caddyHandler := caddyHandlerInit()

	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CaddyHandler: caddyHandler,
	}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {

	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	err = env.Parse(&generalConfig)
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,X-Requested-With")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Origin")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func NewApp() error {
	loadConfig()

	r := gin.Default()

	r.Use(corsMiddleware())
	r.POST("/graphql", graphqlHandler())
	r.GET("/graphql", playgroundHandler())
	r.Run()

	return nil
}
