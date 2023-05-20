package injectors

import (
	"cadigo-api/config"
	"cadigo-api/db/mongodb/infrastructure"
)

func ProvideMongoDBConnector(config *config.Config) (infrastructure.MongodbConnector, error) {
	return infrastructure.NewMongodbConnector(&config.MongodbConfig)
}
