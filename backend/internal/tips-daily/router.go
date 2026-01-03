package tipsdaily

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	// update result match finish
	api.Patch("/match-result", UpdateMatchResult)

}
