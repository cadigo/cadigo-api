package caddy

import (
	"cadigo-api/app/modelapp"
	"cadigo-api/db/mongodb/infrastructure"
	"cadigo-api/db/mongodb/modeldb"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type caddyMongoRepo struct {
	*infrastructure.BaseMongoRepo
	collection string
}

func (repo *caddyMongoRepo) Create(ctx context.Context, record *modelapp.Caddy) (*modelapp.Caddy, error) {
	caddy := modeldb.Caddy{}

	collection := repo.MongodbConnector.DB(ctx).Collection(repo.collection)
	result, err := collection.InsertOne(ctx, caddy)
	if err != nil {
		return record, err
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()
	record.ID = &id

	return record, nil
}

func (repo *caddyMongoRepo) GetByURL(ctx context.Context, url string) (*modeldb.Caddy, error) {
	var (
		err   error
		caddy modeldb.Caddy
	)

	collection := repo.MongodbConnector.DB(ctx).Collection(repo.collection)
	err = collection.FindOne(ctx, bson.M{"url": url}).Decode(&caddy)
	if err != nil {
		return &caddy, err
	}

	return &caddy, nil
}
