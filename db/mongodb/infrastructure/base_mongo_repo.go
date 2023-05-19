package infrastructure

import (
	"context"
	"strconv"
	"strings"
	"time"
)

type MongodbRepositoryCfg struct {
	CustomIDPrefix  string `mapstructure:"custom_id_prefix" json:"custom_id_prefix"`
	EnabledNRTracer bool   `mapstructure:"enabled_newrelic_tracer" json:"enabled_newrelic_tracer"`
}

type BaseMongoRepo struct {
	Config           *MongodbRepositoryCfg
	MongodbConnector MongodbConnector
}

func NewBaseMongoRepo(cfg *MongodbRepositoryCfg, mongodbConnector MongodbConnector) *BaseMongoRepo {
	base := &BaseMongoRepo{
		Config:           cfg,
		MongodbConnector: mongodbConnector,
	}

	return base
}

func (base *BaseMongoRepo) GenerateID(ctx context.Context) string {
	return base.Config.CustomIDPrefix + "-" + strings.ToUpper(strconv.FormatInt(time.Now().UnixNano(), 36))
}
