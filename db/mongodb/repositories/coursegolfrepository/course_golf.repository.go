package coursegolfrepository

import (
	"cadigo-api/app/interface/coursegolfinterface"
	"cadigo-api/app/modela"
	"cadigo-api/db/mongodb/infrastructure"
	"cadigo-api/db/mongodb/modeld"
	"cadigo-api/db/mongodb/repositories"
	"context"

	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	*infrastructure.BaseMongoRepo
	collection string
}

func NewRepository(baseMongoRepo *infrastructure.BaseMongoRepo) coursegolfinterface.CourseGolfRepository {
	return &Repository{
		baseMongoRepo,
		_courseGolfCollection,
	}
}

func (repo *Repository) Create(ctx context.Context, record *modela.CourseGolf) (*modela.CourseGolf, error) {
	courseGolf := new(modeld.CourseGolf).Init()
	err := copier.Copy(&courseGolf, record)

	if err != nil {
		return nil, err
	}

	collection := repo.MongodbConnector.DB(ctx).Collection(repo.collection)
	result, err := collection.InsertOne(ctx, &courseGolf)
	if err != nil {
		return nil, err
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()
	record.ID = id

	return record, nil
}

func (repo *Repository) Update(ctx context.Context, argID string, record *modela.CourseGolf) (*modela.CourseGolf, error) {
	var (
		err    error
		result modeld.CourseGolf
		update bson.M
	)

	courseGolf := new(modeld.CourseGolf).Init()
	err = copier.Copy(&courseGolf, record)
	if err != nil {
		return nil, err
	}

	update, err = repositories.ParseUpdate(courseGolf)
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

	r := result.ToCourseGolf()

	return &r, nil
}

func (repo *Repository) Replace(ctx context.Context, argID string, record *modela.CourseGolf) (*modela.CourseGolf, error) {
	coll := repo.MongodbConnector.DB(ctx).Collection(repo.collection)
	id, _ := primitive.ObjectIDFromHex(argID)
	filter := bson.D{{Key: "_id", Value: id}}
	replacement := modeld.CourseGolf{
		Language: record.Location,
	}

	_, err := coll.ReplaceOne(context.TODO(), filter, replacement)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (repo *Repository) Delete(ctx context.Context, argID string, record *modela.CourseGolf) (*modela.CourseGolf, error) {
	coll := repo.MongodbConnector.DB(ctx).Collection("movies")
	filter := bson.D{{Key: "title", Value: "Twilight"}}
	_, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (repo *Repository) GetByID(ctx context.Context, id string) (*modela.CourseGolf, error) {
	var (
		err        error
		courseGolf modeld.CourseGolf
	)

	collection := repo.MongodbConnector.DB(ctx).Collection(repo.collection)
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = collection.FindOne(ctx, bson.M{"_id": docID}).Decode(&courseGolf)
	if err != nil {
		return nil, err
	}

	c := courseGolf.ToCourseGolf()

	return &c, nil
}

func (repo *Repository) GetAll(ctx context.Context, pagination modela.Pagination) (result []*modela.CourseGolf, total int64, err error) {
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

	courseGolf := []modeld.CourseGolf{}
	if err = curs.All(ctx, &courseGolf); err != nil {
		return result, total, err
	}

	result = []*modela.CourseGolf{}
	for _, c := range courseGolf {
		r := c.ToCourseGolf()
		result = append(result, &r)
	}

	return result, total, err
}
