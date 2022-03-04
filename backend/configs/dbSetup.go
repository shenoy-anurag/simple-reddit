package configs

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateClient() *mongo.Client {
	LoadEnvVariables()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(os.Getenv("MONGO_URI")),
	)
	if err != nil {
		panic(err)
	}
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	return client
}

// MongoDB client instance
var MongoClient *mongo.Client = CreateClient()
var MongoDB *mongo.Database = MongoClient.Database(os.Getenv("DB_NAME"))

// get database collections
func GetCollection(db *mongo.Database, collectionName string) *mongo.Collection {
	collection := db.Collection(collectionName)
	return collection
}
