package mongo

import (
	"context"
	"log"
	"time"

	"writer-api/internal/model"

	"go.mongodb.org/mongo-driver/bson"
)

func SaveStripeEvent(event *model.StripeEvent) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	doc := bson.M{
		"id":       event.ID,
		"type":     event.Type,
		"created":  event.Created,
		"livemode": event.Livemode,
		"data":     event.Data,
		"raw_time": time.Now(),
	}

	_, err := GetCollection().InsertOne(ctx, doc)
	if err != nil {
		log.Printf("Error saving to MongoDB: %v", err)
		return err
	}

	log.Println("Stripe event saved to MongoDB")
	return nil
}
