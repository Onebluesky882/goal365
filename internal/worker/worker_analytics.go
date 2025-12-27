package worker

import (
	m "mytipster/models/fixture"

	"github.com/gofiber/fiber/v2"
)

// AnalyticsFixture handles batch processing of fixture IDs
func AnalyticsFixture(c *fiber.Ctx) error {
	ids := []int{
		1485798,
		// 1486269, 1389669, 1485645, 1394567, 1379114,
		// 	1387830, 1378002, 1389674, 1401121, 1486270, 1387326,
		// 	1488408, 1483270, 1483852, 1451260, 1451432, 1382562,
		// 	1379120, 1378006, 1380388, 1382726, 1394585, 1387351,
		// 	1388439, 1347241, 1489310,
	}
	//id := c.Query("id")

	// transform, err := strconv.Atoi(id)
	// if err != nil {
	// 	return err
	// }

	fixtureBundles, err := HelperSlickFixtureBuddle(ids)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fixtureBundles)
}

// HelperSlickFixtureBuddle processes fixtures concurrently
func HelperSlickFixtureBuddle(id []int) ([]*m.FixturePredictionBundle, error) {
	result := make([]*m.FixturePredictionBundle, 0)

	
}
