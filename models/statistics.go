package models

import "time"

type WorkVolume struct {
	MeterSquare float64   `json:"meter_square"`
	WashedAt    time.Time `json:"washed_at"`
	Type        string    `json:"type"`
}

type WorkVolumeListRequest struct {
	WashedAtFrom time.Time `json:"status,omitempty"`
	WashedAtTo   time.Time `json:"slug,omitempty"`
}
