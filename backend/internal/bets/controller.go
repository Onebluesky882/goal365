package bets

import (
	"fmt"
	analytic_module "mytipster/models/analytic"
	bets_models "mytipster/models/bets"

	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

func InsertPickedHandler(db *bun.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req bets_models.InsertPickedRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid request body",
			})
		}

		ctx := c.Context()

		// เรียก service ตาม signature ใหม่
		if err := InsertPicked(req.Items, req.AnalyticsID, db, ctx); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"success":  true,
			"inserted": len(req.Items),
		})
	}
}

func GetBetListsByDateHandler(db *bun.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		date := c.Query("date")
		ctx := c.Context()
		// fetch analytics from DB
		var analyticsItems []analytic_module.MyAnalytics
		if err := db.NewSelect().Model(&analyticsItems).Scan(ctx); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// 1. Fetch bets for the given date
		bets, err := GetBetListsByDate(date, analyticsItems, db, ctx)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})

		}

		fmt.Println("analyticsItems:", analyticsItems)
		return c.JSON(fiber.Map{
			"success": true,
			"bets":    bets,
			"count":   len(bets),
		})
	}

}

func UpdateMyBetsHandler(db *bun.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body bets_models.Bets
		id := c.Query("id")

		ctx := c.Context()
		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "invalid body",
			})
		}

		err := UpdateMyBets(id, body, db, ctx)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"success": true,
		})
	}

}

