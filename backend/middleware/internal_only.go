package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func InternalOnly(c *fiber.Ctx) error {
	if c.Get("X-Internal-Secret") != os.Getenv("INTERNAL_SECRET") {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	return c.Next()
}
