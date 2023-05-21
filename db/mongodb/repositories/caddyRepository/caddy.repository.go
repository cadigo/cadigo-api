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
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (repo *Repository) Update(ctx context.Context, argID string, record *modelA.Caddy) (*modelA.Caddy, error) {
	coll := repo.MongodbConnector.DB(ctx).Collection(repo.collection)
	id, _ := primitive.ObjectIDFromHex(argID)
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"avg_rating", 4.4}}}}

	result := modelD.Caddy{}
	err := coll.FindOneAndUpdate(
		ctx,
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After), // <- Set option to return document after update (important)
	).Decode(&result)
	if err != nil {
		return nil, err
	}

	r := result.ToCaddy()

	return &r, nil
}

func (repo *Repository) Replace(ctx context.Context, argID string, record *modelA.Caddy) (*modelA.Caddy, error) {
	coll := repo.MongodbConnector.DB(ctx).Collection(repo.collection)
	id, _ := primitive.ObjectIDFromHex(argID)
	filter := bson.D{{"_id", id}}
	replacement := modelD.Caddy{
		Language: record.Location,
	}

	_, err := coll.ReplaceOne(context.TODO(), filter, replacement)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (repo *Repository) Delete(ctx context.Context, argID string, record *modelA.Caddy) (*modelA.Caddy, error) {
	coll := repo.MongodbConnector.DB(ctx).Collection("movies")
	filter := bson.D{{"title", "Twilight"}}
	_, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	return nil, nil
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

func (repo *Repository) GetAll(ctx context.Context, pagination modelA.Pagination) (result []*modelA.Caddy, total int64, err error) {
	query := bson.M{}
	opts := options.Find().
		SetSort(bson.M{pagination.OrderBy: 1}).
		SetLimit(int64(pagination.Limit)).
		SetSkip(int64(pagination.Page * pagination.Limit))

	collection := repo.MongodbConnector.DB(ctx).Collection(repo.collection)
	total, err = collection.CountDocuments(ctx, query)
	if err != nil {
		return result, total, err
	}

	curs, err := collection.Find(ctx, query, opts)
	if err != nil {
		return result, total, err
	}

	caddy := []modelD.Caddy{}
	if err = curs.All(ctx, &caddy); err != nil {
		return result, total, err
	}

	result = []*modelA.Caddy{}
	for _, c := range caddy {
		r := modelA.Caddy{}
		copier.Copy(&r, &c)
		result = append(result, &r)
	}

	return result, total, err
}
