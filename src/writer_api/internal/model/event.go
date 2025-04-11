package model

import (
	"encoding/json"
	"github.com/stripe/stripe-go/v82"
)

func ParseStripeEvent(event stripe.Event) (StripeEvent, error) {
	var setupIntent SetupIntent

	// Deserializar el objeto dentro del evento
	if err := json.Unmarshal(event.Data.Raw, &setupIntent); err != nil {
		return StripeEvent{}, err
	}

	var requestID *string
	if event.Request != nil {
		requestID = stripe.String(event.Request.ID)
	}

	var idempotencyKey *string
	if event.Request != nil {
		idempotencyKey = stripe.String(event.Request.IdempotencyKey)
	}

	return StripeEvent{
		ID:              event.ID,
		Object:          event.Object,
		APIVersion:      event.APIVersion,
		Created:         event.Created,
		Data:            StripeEventData{Object: setupIntent},
		Livemode:        event.Livemode,
		PendingWebhooks: event.PendingWebhooks,
		Request: StripeEventRequest{
			ID:             requestID,
			IdempotencyKey: idempotencyKey,
		},
		Type: string(event.Type),
	}, nil
}

type StripeEvent struct {
	ID              string             `json:"id"`
	Object          string             `json:"object"`
	APIVersion      string             `json:"api_version"`
	Created         int64              `json:"created"`
	Data            StripeEventData    `json:"data"`
	Livemode        bool               `json:"livemode"`
	PendingWebhooks int64              `json:"pending_webhooks"`
	Request         StripeEventRequest `json:"request"`
	Type            string             `json:"type"`
}

type StripeEventData struct {
	Object SetupIntent `json:"object"`
}

type StripeEventRequest struct {
	ID             *string `json:"id"`
	IdempotencyKey *string `json:"idempotency_key"`
}

type SetupIntent struct {
	ID                      string               `json:"id"`
	Object                  string               `json:"object"`
	Application             *string              `json:"application"`
	AutomaticPaymentMethods *string              `json:"automatic_payment_methods"` // puede ser objeto o null, ajusta si necesitas más detalle
	CancellationReason      *string              `json:"cancellation_reason"`
	ClientSecret            string               `json:"client_secret"`
	Created                 int64                `json:"created"`
	Customer                *string              `json:"customer"`
	Description             *string              `json:"description"`
	FlowDirections          *string              `json:"flow_directions"`
	LastSetupError          *string              `json:"last_setup_error"` // puede modelarse como struct si lo necesitas
	LatestAttempt           *string              `json:"latest_attempt"`
	Livemode                bool                 `json:"livemode"`
	Mandate                 *string              `json:"mandate"`
	Metadata                map[string]string    `json:"metadata"`
	NextAction              *string              `json:"next_action"` // puede ser objeto si quieres modelarlo
	OnBehalfOf              *string              `json:"on_behalf_of"`
	PaymentMethod           string               `json:"payment_method"`
	PaymentMethodOptions    PaymentMethodOptions `json:"payment_method_options"`
	PaymentMethodTypes      []string             `json:"payment_method_types"`
	SingleUseMandate        *string              `json:"single_use_mandate"`
	Status                  string               `json:"status"`
	Usage                   string               `json:"usage"`
}

type PaymentMethodOptions struct {
	ACSSDebit *ACSSDebitOptions `json:"acss_debit"`
}

type ACSSDebitOptions struct {
	Currency           string          `json:"currency"`
	MandateOptions     *MandateOptions `json:"mandate_options"`
	VerificationMethod string          `json:"verification_method"`
}

type MandateOptions struct {
	IntervalDescription string `json:"interval_description"`
	PaymentSchedule     string `json:"payment_schedule"`
	TransactionType     string `json:"transaction_type"`
}
