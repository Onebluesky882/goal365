package analytics

import (
	"fmt"
	"log"
	"mytipster/lib"
	m "mytipster/models/analytic"

	"github.com/gofiber/fiber/v2"
)

func WritePredictionsHandler(service AnalyticService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		date := c.Query("date")
		if date == "" {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "require date 2025-12-01",
			})
		}
		oddsMap, err := lib.ReadOddsMap(fmt.Sprintf("bin/%s/filtered_odds.json", date))
		if err != nil {
			log.Printf("❌ ไม่สามารถอ่านไฟล์: %v\n", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "cannot read output.json",
			})
		}

		log.Printf("✅ พบ %d fixtures ใน bin\n", len(oddsMap))

		// สร้าง slice ของ IDs
		ids := make([]string, 0, len(oddsMap))
		for id := range oddsMap {
			ids = append(ids, id)
		}

		// ประมวลผล
		resp, err := PredictionsMany(date, ids, oddsMap)
		if err != nil {
			log.Printf("❌ Error in Predictions: %v\n", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		fileName := "predictions.json"
		if err := lib.WriteJSONWithCustomDate(date, fileName, resp.Items); err != nil {
			log.Println(err)
		} else {
			log.Println("🎉 เขียนไฟล์สำเร็จ:", fileName)
		}

		log.Printf("🎉 ส่งผลลัพธ์ %d รายการ\n", len(resp.Items))
		return c.JSON(fiber.Map{
			"success": true,
			"updated": len(resp.Items),
		})
	}
}

func GetPredictionByDayHandler(service AnalyticService) fiber.Handler {

	return func(c *fiber.Ctx) error {

		ctx := c.Context()
		date := c.Query("date")
		log.Println("date =", date)
		predictions, err := service.PredictionByDay(ctx, date)

		log.Println("predictions :", predictions)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"success":     true,
			"predictions": predictions,
		})
	}
}

func InsertRetryPrediction(service AnalyticService) fiber.Handler {

	return func(c *fiber.Ctx) error {
		ctx := c.Context()
		date := c.Query("date")

		data, err := RetryAndCollectItems(date)
		if err != nil {
			log.Printf("❌ RetryAndCollectItems error: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		if len(data) > 0 {
			if err := service.InsertMany(ctx, data); err != nil {
				log.Printf("❌ Insert to DB failed: %v\n", err)
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "insert to DB failed",
				})
			}
		}

		return c.JSON(fiber.Map{
			"status":   "success",
			"inserted": len(data),
		})
	}
}

func InsertPredictions(service AnalyticService) fiber.Handler {

	return func(c *fiber.Ctx) error {

		ctx := c.Context()
		date := c.Query("date")

		// insert db
		data, err := lib.ReadJson[[]m.MyAnalytics](fmt.Sprintf("bin/%s/predictions.json", date))

		if err != nil {
			log.Fatalf("❌ Cannot read predictions.json: %v", err)
		}

		// insert many
		if err := service.InsertMany(ctx, data); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": fmt.Sprintf("Insert failed: %v", err),
			})
		}
		return c.JSON(fiber.Map{"status": "success"})
	}
}

func naWinTaTipsHandler(service AnalyticService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		fixtureID := c.Query("id")
		if fixtureID == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "missing query parameter: id",
			})
		}

		tip, err := service.naWinTaTips(ctx, fixtureID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("insert failed: %v", err),
			})
		}
		return c.JSON(fiber.Map{
			"status":   "success",
			"inserted": tip,
		})
	}

}
