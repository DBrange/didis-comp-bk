package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	// usr      = "admin"
	// pwd      = "password"
	// host     = "mongo"
	// port     = "27017"
	database = "didis-comp-bk-api"
)

var MongoClient *mongo.Client

func init() {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = "mongodb://admin:password@mongo:27017/didis-comp-bk?authSource=admin"
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	MongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("connection error", err)
	}

	// Retry logic
	for i := 0; i < 5; i++ {
		err = MongoClient.Ping(ctx, readpref.Primary())
		if err == nil {
			break
		}
		log.Printf("Failed to ping MongoDB. Retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatal("Failed to connect to MongoDB after 5 attempts", err)
	}
}

func GetCollection(collection string) *mongo.Collection {
	return MongoClient.Database(database).Collection(collection)
}
