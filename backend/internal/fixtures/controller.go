package fixtures

import (
	"github.com/gofiber/fiber/v2"
)

func GetFixtureDate(c *fiber.Ctx) error {

	date := c.Query("date")
	res, err := QueryFixtureDate(date)
	if err != nil {
		return err
	}

	return c.JSON(res)
}

func GetFixtureById(c *fiber.Ctx) error {

	id := c.Query("id")
	res, err := QueryFixtureId(id)
	if err != nil {
		return err
	}

	return c.JSON(res)
}

func Odds(c *fiber.Ctx) error {
	id := c.Query("id")
	res, err := QueryFixtureOdds(id)
	if err != nil {
		return err
	}

	return c.JSON(res)

}

func Predictions(c *fiber.Ctx) error {
	id := c.Query("id")
	resp, err := QueryPrediction(id)
	if err != nil {
		return err
	}
	return c.JSON(resp)
}

func GetFixtureIds(c *fiber.Ctx) error {
	date := c.Query("date")

	resp, err := GetIds(date)
	if err != nil {
		return err
	}
	return c.JSON(resp)
}
