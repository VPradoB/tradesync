package web_hook

import (
	"log"
	"writer-api/internal/kafka"
	"writer-api/internal/model"
	"writer-api/internal/mongo"
	"writer-api/internal/sqlite"

	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/webhook"
)

func WebHook(stripeWebhookSecret string) *fiber.App {
	app := fiber.New()

	app.Post("webhook", func(c *fiber.Ctx) error {
		signature := c.Get("Stripe-Signature")

		event, err := webhook.ConstructEvent(c.Body(), signature, stripeWebhookSecret)
		if err != nil {
			log.Printf("Error verificando la firma: %v\n", err)
			return c.Status(fiber.StatusBadRequest).SendString("Firma inválida")
		}
		err = processStripeEvent(event)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Firma inválida")
		}
		return c.SendStatus(fiber.StatusOK)
	})

	return app
}

func processStripeEvent(event stripe.Event) error {
	model, err := model.ParseStripeEvent(event)
	if err != nil {
		log.Printf("Error parseando el evento: %v\n", err)
		return err
	}
	err = kafka.SendStripeEvent(model)
	if err != nil {
		log.Printf("Error sending event to kafka")
		sqlite.SaveFailedEvent(model, true, true)
		return err
	}

	err = mongo.SaveStripeEvent(model)
	if err != nil {
		log.Printf("Error saving event in mongo")
		sqlite.SaveFailedEvent(model, true, false)
		return err
	}
	return nil
}
