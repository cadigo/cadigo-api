package config

type Config struct {
	Environment string `env:"ENVIRONMENT,required"`
	Version     int    `env:"VERSION,required"`
}
