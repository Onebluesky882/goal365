package matchresults

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

func RegisterRoutes(app *fiber.App, db *bun.DB) {
	api := app.Group("/api")
	sav := NewMatchResultService(db)

	// update result match finish
	api.Patch("/match-result", MatchResultHandler(sav))

}
