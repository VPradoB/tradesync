package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/webhook"
)

func main() {
	stripe.Key = os.Getenv("STRIPE_API_SECRET")
	stripeWebhookSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")
	app := fiber.New()

	app.Post("/webhook", func(c *fiber.Ctx) error {
		signature := c.Get("Stripe-Signature")

		event, err := webhook.ConstructEvent(c.Body(), signature, stripeWebhookSecret)
		if err != nil {
			log.Printf("Error verificando la firma: %v\n", err)
			return c.Status(fiber.StatusBadRequest).SendString("Firma inválida")
		}

		switch event.Type {
		case "checkout.session.completed":
			log.Println("✅ Pago completado:", event.ID)
			// Aquí puedes actualizar tu base de datos, enviar emails, etc.
		default:
			log.Println("ℹ️ Evento recibido:", event.Type)
		}

		return c.SendStatus(fiber.StatusOK)
	})

	// Inicia el servidor en el puerto 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Println("🚀 Webhook escuchando en http://localhost:" + port)
	log.Fatal(app.Listen(":" + port))
}
