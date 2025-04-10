package mongo

import "go.mongodb.org/mongo-driver/mongo"

func GetCollection() *mongo.Collection {
	dbName := "stripe_events"
	collectionName := "webhooks"

	return GetClient().Database(dbName).Collection(collectionName)
}
