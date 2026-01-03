package mybets

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(c *fiber.App) {
	api := c.Group("/api")
	api.Post("/create", CreateMyBets)
}
