package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	// usr      = "didis-comp-bk"
	// pwd      = "didis-comp-bk"
	// host     = "mongo"
	// port     = "27017"
	database = "didis-comp-bk"
)

var MongoClient *mongo.Client

func init() {
	uri := "mongodb://localhost:27017"

	var err error
	MongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("connection error", err)
	}

	err = MongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("ping failed", err)
	}
}

func GetCollection(collection string) *mongo.Collection {
	return MongoClient.Database(database).Collection(collection)
}
