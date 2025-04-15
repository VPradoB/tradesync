package main

import (
	"log"
	"os"
	web_hook "writer-api/internal/api"
	"writer-api/internal/kafka"
	"writer-api/internal/mongo"
	"writer-api/internal/sqlite"
)

func main() {
	// init kafka
	kafkaBroker := []string{os.Getenv("KAFKA_BROKER_CONNECTION_STRING")}
	err := kafka.InitProducer(kafkaBroker, os.Getenv("KAFKA_TOPIC"))
	if err != nil {
		log.Fatalf("error initializing kafka producer %v", err)
	}
	defer kafka.CloseProducer()

	// init mongo
	err = mongo.InitMongo()
	if err != nil {
		log.Fatalf("error establishing mongodb connection %v", err)
	}

	// init sqlite
	err = sqlite.NewSQLiteStore("./db.sqlite")
	if err != nil {
		log.Fatalf("error with the sqlite store definition %s", err)
	}
	defer sqlite.CloseConnection()

	// init api
	app := web_hook.WebHook(os.Getenv("STRIPE_WEBHOOK_SECRET"))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Fatal(app.Listen(":" + port))
}
