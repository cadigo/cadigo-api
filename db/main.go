package db

import (
	"context"
	"fmt"
	"net/url"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	username := "cadigo"
	password := "psmfuXL05N07dSiZ"
	cluster := "cadigo-dev.nnitvob.mongodb.net"
	// @bistro-property.0f3jw.gcp.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	uri := "mongodb+srv://" + url.QueryEscape(username) + ":" +
		url.QueryEscape(password) + "@" + cluster +
		"/?retryWrites=true&w=majority"

	// fmt.Printf(uri)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("<dbName>").Collection("<collName>")

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var results []bson.D
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}
}
