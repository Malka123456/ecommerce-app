package api

import (
	"ecommerce-app/config"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	app.Get("/health", HealthCheck)

	if err := app.Listen(config.ServerPort); err != nil {
		panic(err)
	}
}

func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "I am breathing!",
	})
}