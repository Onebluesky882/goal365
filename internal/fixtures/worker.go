package fixtures

import (
	"fmt"
	"mytipster/api"
	m "mytipster/models/fixture"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func WorkerFixtureService(c *fiber.Ctx) error {
	date := c.Query("date")

	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/fixtures?date=%s", date)
	resp, err := api.Fixtures("GET", url)
	if err != nil {
		return err
	}

	// step
	// got slice of fixture Ids
	result, err := getFixtureIds(resp)
	if err != nil {
		return err
	}

	// slice call function fixtureOddsPrediction

	fixtureBuddle, err := HelperSlickFixtureBuddle(result)
	if err != nil {
		return err
	}

	return c.JSON(fixtureBuddle)


	// next filter match tips Football with strategy




	// upload big query
}

func HelperSlickFixtureBuddle(ids []int) ([]*m.FixturePredictionBundle, error) {

	/*  เหตุผลที่ใช้ []*Bundle

	•	ไม่ copy struct ใหญ่ ๆ
	•	ทำงานเร็ว
	•	เป็น pattern มาตรฐานใน Go backend

	*/

	result := make([]*m.FixturePredictionBundle, 0, len(ids))
	for _, v := range ids {
		id := strconv.Itoa(v)

		temp, err := fixtureOddsPrediction(id)
		if err != nil {
			return nil, err
		}

		// คือขั้นตอน   wrap to struct
		bundle := &m.FixturePredictionBundle{
			FixtureIDs: []int{v},
			Items:      []m.FixturePrediction{*temp},
		}
		result = append(result, bundle)
	}
	return result, nil
}
