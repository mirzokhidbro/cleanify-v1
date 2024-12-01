package models

import "time"

type StatusChangeHistory struct {
	Status    int8      `json:"status"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateStatusChangeHistoryModel struct {
	HistoryableType string
	HistoryableID   int
	UserID          string
	CompanyID       string
	Status          int8
}
