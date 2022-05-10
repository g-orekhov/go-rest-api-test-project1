package db

import (
	"context"
	"fmt"
)

type Coords struct {
	Long float64 `json:"long"`
	Lat  float64 `json:"lat"`
}

type EventModel struct {
	Id               int      `json:"id"`
	Title            string   `json:"title"`
	ShortDescription string   `json:"shortDescription"`
	Description      string   `json:"description"`
	Coords           Coords   `json:"coords"`
	Images           []string `json:"images"`
	Preview          string   `json:"prewiew"`
}

func (db *DB) CreateEvent(ctx context.Context, m *EventModel) (*EventModel, error) {

	query := fmt.Sprintf("INSERT INTO events (%v)", "df")
	err := db.Postgres.QueryRow(ctx, query).Scan(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
