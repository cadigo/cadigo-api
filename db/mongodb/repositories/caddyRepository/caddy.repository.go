package caddyRepository

import (
	"cadigo-api/app/interface/caddyInterface"
	"cadigo-api/app/modelA"
	"cadigo-api/db/mongodb/infrastructure"
	"cadigo-api/db/mongodb/modelD"
	"context"
	"encoding/json"
	"fmt"

	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository struct {
	*infrastructure.BaseMongoRepo
	collection string
}

func NewRepository(baseMongoRepo *infrastructure.BaseMongoRepo) caddyInterface.CaddyRepository {
	return &Repository{
		baseMongoRepo,
		_caddyeCollection,
	}
}

func (repo *Repository) Create(ctx context.Context, record *modelA.Caddy) (*modelA.Caddy, error) {
	caddy := modelD.Caddy{}
	err := copier.Copy(&caddy, record)

	if err != nil {
		return nil, err
	}

	{
		c, _ := json.Marshal(caddy)
		fmt.Println(string(c))
	}

	collection := repo.MongodbConnector.DB(ctx).Collection(repo.collection)
	result, err := collection.InsertOne(ctx, caddy)
	if err != nil {
		return record, err
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()
	record.ID = id

	return record, nil
}

func (repo *Repository) GetByURL(ctx context.Context, url string) (*modelD.Caddy, error) {
	var (
		err   error
		caddy modelD.Caddy
	)

	collection := repo.MongodbConnector.DB(ctx).Collection(repo.collection)
	err = collection.FindOne(ctx, bson.M{"url": url}).Decode(&caddy)
	if err != nil {
		return &caddy, err
	}

	return &caddy, nil
}
