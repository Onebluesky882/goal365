package bigquery

import (
	"fmt"
	"time"

	"mytipster/api"
	"mytipster/lib/helper"
	m "mytipster/models/fixture"

	"github.com/gofiber/fiber/v2"
)

func Service(c *fiber.Ctx) error {
	date := c.Query("date")
	if date == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "date is required (format: YYYY-MM-DD)",
		})
	}

	url := fmt.Sprintf(
		"https://api-football-v1.p.rapidapi.com/v3/fixtures?date=%s",
		date,
	)

	resp, err := api.Fixtures("GET", url)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if len(resp.Response) == 0 {
		return c.JSON(fiber.Map{
			"success": true,
			"count":   0,
			"message": "No fixtures found for this date",
		})
	}

	// แปลง Response เป็น BigQuery rows
	var rows []m.FixtureBigQuery

	for _, f := range resp.Response {
		row := m.FixtureBigQuery{
			IngestionTime: time.Now(), // เปลี่ยนจาก PartitionTime

			// Fixture
			Fixture: m.FixtureBQ{
				ID:        int64(f.Fixture.ID),
				Referee:   helper.ToNullString(f.Fixture.Referee),
				Timezone:  f.Fixture.Timezone,
				Date:      f.Fixture.Date,
				Timestamp: f.Fixture.Timestamp,
				Periods: m.PeriodsBQ{
					First:  helper.ToNullInt64(f.Fixture.Periods.First),
					Second: helper.ToNullInt64(f.Fixture.Periods.Second),
				},
				Venue: m.VenueBQ{
					ID:   helper.ToNullInt64(f.Fixture.Venue.ID),
					Name: helper.ToNullString(f.Fixture.Venue.Name),
					City: helper.ToNullString(f.Fixture.Venue.City),
				},
				Status: m.StatusBQ{
					Long:    f.Fixture.Status.Long,
					Short:   f.Fixture.Status.Short,
					Elapsed: helper.ToNullInt64(f.Fixture.Status.Elapsed),
					Extra:   helper.ToNullInt64(f.Fixture.Status.Extra),
				},
			},

			// League
			League: m.LeagueBQ{
				ID:        int64(f.League.ID),
				Name:      f.League.Name,
				Country:   f.League.Country,
				Logo:      f.League.Logo,
				Flag:      helper.ToNullString(f.League.Flag),
				Season:    int64(f.League.Season),
				Round:     f.League.Round,
				Standings: f.League.Standings,
			},

			// Teams
			Teams: m.TeamsBQ{
				Home: m.TeamBQ{
					ID:     int64(f.Teams.Home.ID),
					Name:   f.Teams.Home.Name,
					Logo:   f.Teams.Home.Logo,
					Winner: helper.ToNullBool(f.Teams.Home.Winner),
				},
				Away: m.TeamBQ{
					ID:     int64(f.Teams.Away.ID),
					Name:   f.Teams.Away.Name,
					Logo:   f.Teams.Away.Logo,
					Winner: helper.ToNullBool(f.Teams.Away.Winner),
				},
			},

			// Goals
			Goals: m.GoalsBQ{
				Home: helper.ToNullInt64(f.Goals.Home),
				Away: helper.ToNullInt64(f.Goals.Away),
			},

			// Score
			Score: m.ScoreBQ{
				Halftime: m.ScoreDetailBQ{
					Home: helper.ToNullInt64(f.Score.Halftime.Home),
					Away: helper.ToNullInt64(f.Score.Halftime.Away),
				},
				Fulltime: m.ScoreDetailBQ{
					Home: helper.ToNullInt64(f.Score.Fulltime.Home),
					Away: helper.ToNullInt64(f.Score.Fulltime.Away),
				},
				Extratime: m.ScoreDetailBQ{
					Home: helper.ToNullInt64(f.Score.Extratime.Home),
					Away: helper.ToNullInt64(f.Score.Extratime.Away),
				},
				Penalty: m.ScoreDetailBQ{
					Home: helper.ToNullInt64(f.Score.Penalty.Home),
					Away: helper.ToNullInt64(f.Score.Penalty.Away),
				},
			},
		}

		rows = append(rows, row)
	}

	// Insert ข้อมูลทั้งหมดลง BigQuery
	if len(rows) > 0 {
		fmt.Printf("📤 Inserting %d fixtures helper.To BigQuery...\n", len(rows))
		Bigquery(rows)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"count":   len(rows),
		"date":    date,
		"message": fmt.Sprintf("Successfully inserted %d fixtures", len(rows)),
	})
}
