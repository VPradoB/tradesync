package kafka

import (
	"encoding/json"
	"log"
	"time"

	"github.com/IBM/sarama"
	"writer-api/internal/model"
)

var producer sarama.SyncProducer
var topic string

// InitProducer crea e inicializa el Kafka SyncProducer.
func InitProducer(brokers []string, t string) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second

	p, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return err
	}

	producer = p
	topic = t
	log.Println("✅ Kafka producer initialized")
	return nil
}

// CloseProducer cierra el producer cuando la app finaliza.
func CloseProducer() {
	if producer != nil {
		if err := producer.Close(); err != nil {
			log.Printf("❌ Error closing Kafka producer: %v", err)
		}
	}
}

// SendStripeEvent serializa y envía el evento al topic.
func SendStripeEvent(event model.StripeEvent) error {
	if producer == nil {
		return sarama.ErrClosedClient
	}

	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(payload),
	}

	_, _, err = producer.SendMessage(msg)
	return err
}
