package config

import "cadigo-api/db/mongodb/infrastructure"

type Config struct {
	Environment string `env:"ENVIRONMENT,required"`
	Version     string `env:"VERSION,required"`
	RedisAddr   string `env:"REDIS_ADDR,required"`
	RedisPass   string `env:"REDIS_PASS,required"`
	infrastructure.MongodbConfig
	infrastructure.MongodbRepositoryCfg
}
