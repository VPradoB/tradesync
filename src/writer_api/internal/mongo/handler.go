package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"writer-api/internal/model"
)

func SaveStripeEvent(event model.StripeEvent) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	doc, err := bson.Marshal(event)
	if err != nil {
		log.Printf("Error marshalling event to BSON: %v", err)
		return err
	}

	var bsonDoc bson.M
	if err := bson.Unmarshal(doc, &bsonDoc); err != nil {
		log.Printf("Error unmarshalling to bson.M: %v", err)
		return err
	}

	// Agregar el raw_time por separado
	bsonDoc["raw_time"] = time.Now()

	_, err = GetCollection().InsertOne(ctx, doc)
	if err != nil {
		log.Printf("Error saving to MongoDB: %v", err)
		return err
	}

	log.Println("Stripe event saved to MongoDB")
	return nil
}
