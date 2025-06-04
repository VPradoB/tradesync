package mongo

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func InitMongo() error {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}

	log.Printf("🔗 Attempting to connect to MongoDB with URI: %s", uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Printf("❌ Error creating MongoDB client: %v", err)
		return err
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Printf("❌ Error pinging MongoDB: %v", err)
		return err
	}

	log.Println("✅ MongoDB connected successfully")
	return nil
}

func GetClient() *mongo.Client {
	return client
}
