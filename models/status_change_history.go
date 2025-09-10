package models

import "time"

type StatusChangeHistory struct {
	Status    int8      `json:"status"`
	Fullname  string    `json:"fullname"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateStatusChangeHistoryModel struct {
	HistoryableType string
	HistoryableID   int
	UserID          int64
	CompanyID       string
	Status          int8
}
