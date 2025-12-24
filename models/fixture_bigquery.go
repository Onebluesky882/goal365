package models

import (
	"time"

	"cloud.google.com/go/bigquery"
)

// FixtureBigQuery - โครงสร้างสำหรับ insert ลง BigQuery
type FixtureBigQuery struct {
	IngestionTime time.Time `bigquery:"ingestion_time"` // เปลี่ยนตรงนี้
	
	Fixture FixtureBQ `bigquery:"fixture"`
	League  LeagueBQ  `bigquery:"league"`
	Teams   TeamsBQ   `bigquery:"teams"`
	Goals   GoalsBQ   `bigquery:"goals"`
	Score   ScoreBQ   `bigquery:"score"`
}

// Fixture Structs
type FixtureBQ struct {
	ID        int64              `bigquery:"id"`
	Referee   bigquery.NullString `bigquery:"referee"`
	Timezone  string             `bigquery:"timezone"`
	Date      string             `bigquery:"date"`
	Timestamp int64              `bigquery:"timestamp"`
	Periods   PeriodsBQ          `bigquery:"periods"`
	Venue     VenueBQ            `bigquery:"venue"`
	Status    StatusBQ           `bigquery:"status"`
}

type PeriodsBQ struct {
	First  bigquery.NullInt64 `bigquery:"first"`
	Second bigquery.NullInt64 `bigquery:"second"`
}

type VenueBQ struct {
	ID   bigquery.NullInt64  `bigquery:"id"`
	Name bigquery.NullString `bigquery:"name"`
	City bigquery.NullString `bigquery:"city"`
}

type StatusBQ struct {
	Long    string             `bigquery:"long"`
	Short   string             `bigquery:"short"`
	Elapsed bigquery.NullInt64 `bigquery:"elapsed"`
	Extra   bigquery.NullInt64 `bigquery:"extra"`
}

// League Struct
type LeagueBQ struct {
	ID        int64               `bigquery:"id"`
	Name      string              `bigquery:"name"`
	Country   string              `bigquery:"country"`
	Logo      string              `bigquery:"logo"`
	Flag      bigquery.NullString `bigquery:"flag"`
	Season    int64               `bigquery:"season"`
	Round     string              `bigquery:"round"`
	Standings bool                `bigquery:"standings"`
}

// Teams Structs
type TeamsBQ struct {
	Home TeamBQ `bigquery:"home"`
	Away TeamBQ `bigquery:"away"`
}

type TeamBQ struct {
	ID     int64              `bigquery:"id"`
	Name   string             `bigquery:"name"`
	Logo   string             `bigquery:"logo"`
	Winner bigquery.NullBool  `bigquery:"winner"`
}

// Goals Struct
type GoalsBQ struct {
	Home bigquery.NullInt64 `bigquery:"home"`
	Away bigquery.NullInt64 `bigquery:"away"`
}

// Score Structs
type ScoreBQ struct {
	Halftime  ScoreDetailBQ `bigquery:"halftime"`
	Fulltime  ScoreDetailBQ `bigquery:"fulltime"`
	Extratime ScoreDetailBQ `bigquery:"extratime"`
	Penalty   ScoreDetailBQ `bigquery:"penalty"`
}

type ScoreDetailBQ struct {
	Home bigquery.NullInt64 `bigquery:"home"`
	Away bigquery.NullInt64 `bigquery:"away"`
}