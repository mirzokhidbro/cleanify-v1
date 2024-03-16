package models

import "time"

type TelegramGroup struct {
	ID                   int
	CompanyID            string
	Name                 string
	NotificationStatuses []int8
	Code                 int
	ChatID               int
	WithLocation         bool
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type CreateTelegramGroupRequest struct {
	ChatID int
	Code   int
	Name   string
}
