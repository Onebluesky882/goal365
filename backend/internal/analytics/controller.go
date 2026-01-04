package analytics

import (
	"context"
	"fmt"
	"log"
	db "mytipster/internal/database"
	"mytipster/lib"
	m "mytipster/models/analytic"

	"github.com/gofiber/fiber/v2"
)

func writePredictions(c *fiber.Ctx) error {
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

func GetPredictionByDay(c *fiber.Ctx) error {
	ctx := context.Background()
	db := db.WithContext(ctx)
	date := c.Query("date")

	predictions, err := PredictionByDay(date, db, ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	fmt.Println("predictions :", predictions)
	return c.JSON(predictions)
}

func insertRetryPrediction(c *fiber.Ctx) error {
	ctx := context.Background()
	db := db.WithContext(ctx)
	date := c.Query("date")

	items, err := RetryAndCollectItems(date)
	if err != nil {
		return err
	}

	if len(items) > 0 {
		if err := insertMany(items, db, ctx); err != nil {
			log.Printf("❌ Insert to DB failed: %v\n", err)
		}
	}

	return c.JSON(fiber.Map{
		"status":   "success",
		"inserted": len(items),
	})
}

func InsertPredictions(c *fiber.Ctx) error {
	ctx := context.Background()
	db := db.WithContext(ctx)
	date := c.Query("date")

	// insert db
	data, err := lib.ReadJson[[]m.MyAnalytics](fmt.Sprintf("bin/%s/predictions.json", date))

	if err != nil {
		log.Fatalf("❌ Cannot read predictions.json: %v", err)
	}

	// insert many
	if err := insertMany(data, db, ctx); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": fmt.Sprintf("Insert failed: %v", err),
		})
	}
	return c.JSON(fiber.Map{"status": "success"})
}
