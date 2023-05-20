package config

import "cadigo-api/db/mongodb/infrastructure"

type Config struct {
	Environment string `env:"ENVIRONMENT,required"`
	Version     int    `env:"VERSION,required"`
	infrastructure.MongodbConfig
	infrastructure.MongodbRepositoryCfg
}
