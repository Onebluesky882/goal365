package fixtures

import (
	"fmt"
	"time"

	"mytipster/api"
	bqservice "mytipster/internal/bigquery"
	"mytipster/models"

	"cloud.google.com/go/bigquery"
	"github.com/gofiber/fiber/v2"
)

// Helper: แปลง *int เป็น NullInt64
func toNullInt64(v *int) bigquery.NullInt64 {
	if v == nil {
		return bigquery.NullInt64{Valid: false}
	}
	return bigquery.NullInt64{Int64: int64(*v), Valid: true}
}

// Helper: แปลง *string เป็น NullString
func toNullString(v *string) bigquery.NullString {
	if v == nil || *v == "" {
		return bigquery.NullString{Valid: false}
	}
	return bigquery.NullString{StringVal: *v, Valid: true}
}

// Helper: แปลง *bool เป็น NullBool
func toNullBool(v *bool) bigquery.NullBool {
	if v == nil {
		return bigquery.NullBool{Valid: false}
	}
	return bigquery.NullBool{Bool: *v, Valid: true}
}

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
	var rows []models.FixtureBigQuery

	for _, f := range resp.Response {
		row := models.FixtureBigQuery{
			IngestionTime: time.Now(), // เปลี่ยนจาก PartitionTime
			
			// Fixture
			Fixture: models.FixtureBQ{
				ID:        int64(f.Fixture.ID),
				Referee:   toNullString(f.Fixture.Referee),
				Timezone:  f.Fixture.Timezone,
				Date:      f.Fixture.Date,
				Timestamp: f.Fixture.Timestamp,
				Periods: models.PeriodsBQ{
					First:  toNullInt64(f.Fixture.Periods.First),
					Second: toNullInt64(f.Fixture.Periods.Second),
				},
				Venue: models.VenueBQ{
					ID:   toNullInt64(f.Fixture.Venue.ID),
					Name: toNullString(f.Fixture.Venue.Name),
					City: toNullString(f.Fixture.Venue.City),
				},
				Status: models.StatusBQ{
					Long:    f.Fixture.Status.Long,
					Short:   f.Fixture.Status.Short,
					Elapsed: toNullInt64(f.Fixture.Status.Elapsed),
					Extra:   toNullInt64(f.Fixture.Status.Extra),
				},
			},
			
			// League
			League: models.LeagueBQ{
				ID:        int64(f.League.ID),
				Name:      f.League.Name,
				Country:   f.League.Country,
				Logo:      f.League.Logo,
				Flag:      toNullString(f.League.Flag),
				Season:    int64(f.League.Season),
				Round:     f.League.Round,
				Standings: f.League.Standings,
			},
			
			// Teams
			Teams: models.TeamsBQ{
				Home: models.TeamBQ{
					ID:     int64(f.Teams.Home.ID),
					Name:   f.Teams.Home.Name,
					Logo:   f.Teams.Home.Logo,
					Winner: toNullBool(f.Teams.Home.Winner),
				},
				Away: models.TeamBQ{
					ID:     int64(f.Teams.Away.ID),
					Name:   f.Teams.Away.Name,
					Logo:   f.Teams.Away.Logo,
					Winner: toNullBool(f.Teams.Away.Winner),
				},
			},
			
			// Goals
			Goals: models.GoalsBQ{
				Home: toNullInt64(f.Goals.Home),
				Away: toNullInt64(f.Goals.Away),
			},
			
			// Score
			Score: models.ScoreBQ{
				Halftime: models.ScoreDetailBQ{
					Home: toNullInt64(f.Score.Halftime.Home),
					Away: toNullInt64(f.Score.Halftime.Away),
				},
				Fulltime: models.ScoreDetailBQ{
					Home: toNullInt64(f.Score.Fulltime.Home),
					Away: toNullInt64(f.Score.Fulltime.Away),
				},
				Extratime: models.ScoreDetailBQ{
					Home: toNullInt64(f.Score.Extratime.Home),
					Away: toNullInt64(f.Score.Extratime.Away),
				},
				Penalty: models.ScoreDetailBQ{
					Home: toNullInt64(f.Score.Penalty.Home),
					Away: toNullInt64(f.Score.Penalty.Away),
				},
			},
		}

		rows = append(rows, row)
	}

	// Insert ข้อมูลทั้งหมดลง BigQuery
	if len(rows) > 0 {
		fmt.Printf("📤 Inserting %d fixtures to BigQuery...\n", len(rows))
		bqservice.Service(rows)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"count":   len(rows),
		"date":    date,
		"message": fmt.Sprintf("Successfully inserted %d fixtures", len(rows)),
	})
}