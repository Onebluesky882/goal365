package mytips

import (
	"fmt"
	"log"
	"mytipster/lib"
	mytips_module "mytipster/models/mytips"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetPredictionByDay(c *fiber.Ctx) error {

	date := c.Query("date")

	predictions, err := PredictionByDay(date)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	fmt.Println("predictions :", predictions)
	return c.JSON(predictions)
}

func InsertPredictions(c *fiber.Ctx) error {
	date := time.Now().Format("2006-01-02")

	// insert db
	data, err := lib.ReadJson[[]mytips_module.MyTipsAnalytics](fmt.Sprintf("bin/%s/predictions.json", date))

	if err != nil {
		log.Fatalf("❌ Cannot read predictions.json: %v", err)
	}

	// insert many
	if err := InsertMany(data); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": fmt.Sprintf("Insert failed: %v", err),
		})
	}
	return c.JSON(fiber.Map{"status": "success"})
}

func WritePrediction(c *fiber.Ctx) error {
	date := time.Now().Format("2006-01-02")
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
	resp, err := Predictions(ids, oddsMap)
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
	return c.JSON(resp.Items)
}

func UpdateMatchResult(c *fiber.Ctx) error {
	// func MatchResult(date string) ([]m.UpdateFixtureResultDTO, error) {
	date := c.Query("date")
	results, err := MatchResult(date)

	if err != nil {
		log.Println("❌ MatchResult error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if len(results) == 0 {
		return c.JSON(fiber.Map{
			"success": true,
			"updated": 0,
			"message": "no fixtures to update",
		})
	}

	if err := UpdateFixtureResult(results); err != nil {
		log.Println("❌ UpdateFixtureResult error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"updated": len(results),
	})
}
