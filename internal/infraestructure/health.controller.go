package infraestructure

import (
	"github.com/gofiber/fiber/v2"
)

func CheckHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("HealthCheck OK")
	}
}
