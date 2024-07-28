package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/DBrange/didis-comp-bk/config"
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
	// uri := os.Getenv("MONGO_URI")
	// if uri == "" {
	// 	uri = "mongodb://admin:password@mongo:27017/didis-comp-bk?authSource=admin"
	// }
	// uri := "mongodb://localhost:27017/"
	// uri := "mongodb://localhost:27018,localhost:27019,localhost:27020/admin?authSource=admin&replicaSet=didi"
	uri := config.Envs.MDBURI
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

	// Initialize the Replica Set
	// err = initiateReplicaSet(ctx, MongoClient)
	// if err != nil {
	// 	log.Fatal("Failed to initiate replica set", err)
	// }
}

func GetCollection(collection string) *mongo.Collection {
	return MongoClient.Database(database).Collection(collection)
}

func GetCollections(collectionNames []string) (map[string]*mongo.Collection, error) {
	collectionMap := make(map[string]*mongo.Collection)

	for _, collName := range collectionNames {
		coll := GetCollection(collName)
		if coll == nil {
			return nil, fmt.Errorf("failed to get collection: %s", collName)
		}
		collectionMap[collName] = coll
	}

	return collectionMap, nil
}

// func initiateReplicaSet(ctx context.Context, client *mongo.Client) error {
// 	var result bson.M
// 	err := client.Database("admin").RunCommand(ctx, bson.D{
// 		{Key: "replSetInitiate", Value: bson.D{
// 			{Key: "_id", Value: "rs0"},
// 			{Key: "members", Value: bson.A{
// 				bson.D{{Key: "_id", Value: 0}, {Key: "host", Value: "localhost:27017"}},
// 			}},
// 		}},
// 	}).Decode(&result)
// 	if err != nil {
// 		return fmt.Errorf("failed to initiate replica set: %v", err)
// 	}
// 	log.Println("Replica set initiated: ", result)
// 	return nil
// }
