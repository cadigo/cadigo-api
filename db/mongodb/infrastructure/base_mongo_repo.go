package infrastructure

import (
	"context"
	"strconv"
	"strings"
	"time"
)

type MongodbRepositoryCfg struct {
	CustomIDPrefix string `env:"CUSTOM_ID_PREFIX,required" json:"custom_id_prefix"`
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
