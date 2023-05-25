package injectors

import (
	"cadigo-api/config"
	"cadigo-api/db/mongodb/infrastructure"
)

func ProvideBaseMongoRepo(config *config.Config,
	mongodbConnector infrastructure.MongodbConnector) *infrastructure.BaseMongoRepo {
	return infrastructure.NewBaseMongoRepo(&config.MongodbRepositoryCfg, mongodbConnector)
}
