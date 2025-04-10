package mongo_test

import (
	"os"
	"testing"

	"writer-api/internal/model"
	"writer-api/internal/mongo"
)

func TestSaveStripeEvent(t *testing.T) {
	_ = os.Setenv("MONGODB_URI", "mongodb://root:example@localhost:27017")

	if err := mongo.InitMongo(); err != nil {
		t.Fatalf("❌ Failed to init Mongo: %v", err)
	}

	event := &model.StripeEvent{
		ID:       "evt_test_123",
		Object:   "event",
		Type:     "setup_intent.created",
		Created:  1686089970,
		Livemode: false,
		Data: model.StripeEventData{
			Object: model.SetupIntent{
				ID:                 "seti_test_123",
				Object:             "setup_intent",
				ClientSecret:       "secret_123",
				Created:            1686089970,
				PaymentMethod:      "pm_test_123",
				PaymentMethodTypes: []string{"acss_debit"},
				Status:             "requires_confirmation",
				Usage:              "off_session",
			},
		},
	}

	if err := mongo.SaveStripeEvent(event); err != nil {
		t.Errorf("❌ Failed to save Stripe event: %v", err)
	} else {
		t.Logf("✅ Stripe event saved successfully")
	}
}
