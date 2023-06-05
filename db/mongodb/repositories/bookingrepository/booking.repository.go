package bookingrepository

import (
	"cadigo-api/app/interface/bookinginterface"
	"cadigo-api/app/modela"
	"cadigo-api/db/mongodb/infrastructure"
	"cadigo-api/db/mongodb/modeld"
	"cadigo-api/db/mongodb/repositories"
	"context"
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

func NewRepository(baseMongoRepo *infrastructure.BaseMongoRepo) bookinginterface.BookingRepository {
	return &Repository{
		baseMongoRepo,
		_bookingCollection,
	}
}

func (repo *Repository) Create(ctx context.Context, record *modela.Booking) (*modela.Booking, error) {
	if record == nil {
		return nil, fmt.Errorf("record is nil")
	}

	booking, err := new(modeld.Booking).Init().SetBooking(*record)
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&booking, record)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	collection := repo.MongodbConnector.DB(ctx).Collection(repo.collection)
	result, err := collection.InsertOne(ctx, &booking)
	if err != nil {
		return nil, err
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()
	record.ID = id

	return record, nil
}

func (repo *Repository) Update(ctx context.Context, argID string, record *modela.Booking) (*modela.Booking, error) {
	var (
		err    error
		result modeld.Booking
		update bson.M
	)

	booking := new(modeld.Booking).Init()
	err = copier.Copy(&booking, record)
	if err != nil {
		return nil, err
	}

	update, err = repositories.ParseUpdate(booking)
	if err != nil {
		return nil, err
	}

	coll := repo.MongodbConnector.DB(ctx).Collection(repo.collection)
	id, _ := primitive.ObjectIDFromHex(argID)
	filter := bson.D{{Key: "_id", Value: id}}

	err = coll.FindOneAndUpdate(
		ctx,
		filter,
		bson.D{{Key: "$set", Value: update}},
		options.FindOneAndUpdate().SetReturnDocument(options.After), // <- Set option to return document after update (important)
	).Decode(&result)
	if err != nil {
		return nil, err
	}

	r := result.ToBooking()

	return &r, nil
}

func (repo *Repository) Replace(ctx context.Context, argID string, record *modela.Booking) (*modela.Booking, error) {
	coll := repo.MongodbConnector.DB(ctx).Collection(repo.collection)
	id, _ := primitive.ObjectIDFromHex(argID)
	filter := bson.D{{Key: "_id", Value: id}}
	replacement := modeld.Booking{
		// Language: record.Location,
	}

	_, err := coll.ReplaceOne(context.TODO(), filter, replacement)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (repo *Repository) Delete(ctx context.Context, argID string, record *modela.Booking) (*modela.Booking, error) {
	coll := repo.MongodbConnector.DB(ctx).Collection("movies")
	filter := bson.D{{Key: "title", Value: "Twilight"}}
	_, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (repo *Repository) GetByID(ctx context.Context, id string) (result *modela.Booking, err error) {
	var (
		booking modeld.Booking
	)

	collection := repo.MongodbConnector.DB(ctx).Collection(repo.collection)
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = collection.FindOne(ctx, bson.M{"_id": docID}).Decode(&booking)
	if err != nil {
		return nil, err
	}

	b := booking.ToBooking()

	return &b, nil
}

func (repo *Repository) GetAll(ctx context.Context, pagination modela.Pagination) (result []*modela.Booking, total int64, err error) {
	query := bson.M{}
	opts := options.Find().
		SetSort(bson.M{pagination.OrderBy: 1}).
		SetLimit(int64(pagination.Limit)).
		SetSkip(int64((pagination.Page - 1) * pagination.Limit))

	collection := repo.MongodbConnector.DB(ctx).Collection(repo.collection)
	total, err = collection.CountDocuments(ctx, query)
	if err != nil {
		return result, total, err
	}

	curs, err := collection.Find(ctx, query, opts)
	if err != nil {
		return result, total, err
	}

	booking := []modeld.Booking{}
	if err = curs.All(ctx, &booking); err != nil {
		return result, total, err
	}

	result = []*modela.Booking{}
	for _, c := range booking {
		r := c.ToBooking()
		result = append(result, &r)
	}

	return result, total, err
}
