package mytips

import (
	"fmt"
	"log"
	mytips_db "mytipster/internal/db/mytips"
	"mytipster/lib"
	mytips_module "mytipster/models/mytips"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Insert(c *fiber.Ctx) error {
	date := time.Now().Format("2006-01-02")

	// insert db
	data, err := lib.ReadJson[[]mytips_module.MyTipsAnalytics](fmt.Sprintf("bin/%s/predictions.json", date))

	if err != nil {
		log.Fatalf("❌ Cannot read predictions.json: %v", err)
	}

	// insert many
	if err := mytips_db.InsertMany(data); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": fmt.Sprintf("Insert failed: %v", err),
		})
	}
	return c.JSON(fiber.Map{"status": "success"})
}
