package models

import (
	"time"
)

type RootFixtureResponse struct {
	Get        string            `json:"get"`
	Parameters FixtureParameters `json:"parameters"`
	Errors     any               `json:"errors"`
	Results    int               `json:"results"`
	Paging     FixturePaging     `json:"paging"`
	Response   []FixtureResponse `json:"response"`
	Info       string            `json:"info"`
}

type FixtureParameters struct {
	Date string `json:"date"`
}

type FixturePaging struct {
	Current int `json:"current"`
	Total   int `json:"total"`
}

type FixtureResponse struct {
	PartitionTime time.Time   `bigquery:"_PARTITIONTIME"`
	FixtureInfo   FixtureInfo `json:"fixture_info" bigquery:"fixture_info"`
	League        League      `json:"league" bigquery:"league"`
	Teams         Teams       `json:"teams" bigquery:"teams"`
	Goals         Goals       `json:"goals" bigquery:"goals"`
	Score         Score       `json:"score" bigquery:"score"`
}
type FixtureInfo struct {
	ID        int     `json:"id"`
	Referee   *string `json:"referee,omitempty"`
	Timezone  string  `json:"timezone"`
	Date      string  `json:"date"`
	Timestamp int64   `json:"timestamp"`
	Periods   Periods `json:"periods"`
	Venue     Venue   `json:"venue"`
	Status    Status  `json:"status"`
}

type Periods struct {
	First  *int `json:"first,omitempty"`
	Second *int `json:"second,omitempty"`
}

type Venue struct {
	ID   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	City *string `json:"city,omitempty"`
}

type Status struct {
	Long    string `json:"long"`
	Short   string `json:"short"`
	Elapsed *int   `json:"elapsed,omitempty"`
	Extra   *int   `json:"extra,omitempty"`
}

type FixtureLeague struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Country   string  `json:"country"`
	Logo      string  `json:"logo"`
	Flag      *string `json:"flag,omitempty"`
	Season    int     `json:"season"`
	Round     string  `json:"round"`
	Standings bool    `json:"standings"`
}

type Teams struct {
	Home Home `json:"home"`
	Away Away `json:"away"`
}

type Home struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Logo   string `json:"logo"`
	Winner *bool  `json:"winner,omitempty"`
}

type Away struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Logo   string `json:"logo"`
	Winner *bool  `json:"winner,omitempty"`
}

type Goals struct {
	Home *int `json:"home,omitempty"`
	Away *int `json:"away,omitempty"`
}

type Score struct {
	Halftime  Halftime  `json:"halftime"`
	Fulltime  Fulltime  `json:"fulltime"`
	Extratime Extratime `json:"extratime"`
	Penalty   Penalty   `json:"penalty"`
}

type Halftime struct {
	Home *int `json:"home,omitempty"`
	Away *int `json:"away,omitempty"`
}

type Fulltime struct {
	Home *int `json:"home,omitempty"`
	Away *int `json:"away,omitempty"`
}

type Extratime struct {
	Home *int `json:"home,omitempty"`
	Away *int `json:"away,omitempty"`
}

type Penalty struct {
	Home *int `json:"home,omitempty"`
	Away *int `json:"away,omitempty"`
}
