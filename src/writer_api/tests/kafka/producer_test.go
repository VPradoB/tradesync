package kafka_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"writer-api/internal/kafka"
	"writer-api/tests/model"
)

func TestSendStripeEvent(t *testing.T) {
	brokers := []string{"localhost:29092"}
	kafka.InitProducer(brokers)
	defer kafka.CloseProducer()
	event := mock_event.BuildMockStripeEvent()
	err := kafka.SendStripeEvent("stripe-events", event)
	assert.NoError(t, err)
	fmt.Println("event sended.")
}
