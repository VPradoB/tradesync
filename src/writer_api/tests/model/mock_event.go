package mock_event

import (
	"time"
	"writer-api/internal/model"
)

func BuildMockStripeEvent() model.StripeEvent {
	now := time.Now().Unix()
	app := "app_test_123"
	autoPay := "enabled"
	cancelReason := "abandoned"
	customer := "cus_test_456"
	description := "Setup for off-session usage"
	flowDir := "inbound"
	lastErr := "none"
	latestAttempt := "attempt_123"
	mandate := "mandate_001"
	nextAction := "confirm"
	onBehalf := "acct_789"
	singleUse := "sum_001"
	reqID := "req_123"
	idemKey := "idem_456"

	return model.StripeEvent{
		ID:              "evt_test_123",
		Object:          "event",
		APIVersion:      "2022-11-15",
		Created:         now,
		Livemode:        false,
		PendingWebhooks: 1,
		Request: model.StripeEventRequest{
			ID:             &reqID,
			IdempotencyKey: &idemKey,
		},
		Type: "setup_intent.created",
		Data: model.StripeEventData{
			Object: model.SetupIntent{
				ID:                      "seti_test_456",
				Object:                  "setup_intent",
				Application:             &app,
				AutomaticPaymentMethods: &autoPay,
				CancellationReason:      &cancelReason,
				ClientSecret:            "seti_secret_test_abc",
				Created:                 now,
				Customer:                &customer,
				Description:             &description,
				FlowDirections:          &flowDir,
				LastSetupError:          &lastErr,
				LatestAttempt:           &latestAttempt,
				Livemode:                false,
				Mandate:                 &mandate,
				Metadata: map[string]string{
					"source": "test",
				},
				NextAction:    &nextAction,
				OnBehalfOf:    &onBehalf,
				PaymentMethod: "pm_test_789",
				PaymentMethodOptions: model.PaymentMethodOptions{
					ACSSDebit: &model.ACSSDebitOptions{
						Currency: "cad",
						MandateOptions: &model.MandateOptions{
							IntervalDescription: "First day of every month",
							PaymentSchedule:     "interval",
							TransactionType:     "personal",
						},
						VerificationMethod: "automatic",
					},
				},
				PaymentMethodTypes: []string{"acss_debit"},
				SingleUseMandate:   &singleUse,
				Status:             "requires_confirmation",
				Usage:              "off_session",
			},
		},
	}
}
