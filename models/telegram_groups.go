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

type TelegramGroupGetListResponse struct {
	ID                   int
	CompanyID            string
	Name                 string
	NotificationStatuses *[]interface{}
	WithLocation         *bool
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type TelegramGroupGetByPrimayKeyResponse struct {
	ID                   int            `json:"id"`
	Name                 string         `json:"name"`
	NotificationStatuses *[]interface{} `json:"notification_statuses"`
	WithLocation         *bool          `json:"with_location"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
}

type CreateTelegramGroupRequest struct {
	ChatID int
	Code   int
	Name   string
}

type GroupVerificationRequest struct {
	Code      int    `json:"code" binding:"required"`
	CompanyID string `json:"company_id" binding:"required"`
}

type TelegramGroupEditRequest struct {
	WithLocation         *bool  `json:"with_location" binding:"required"`
	NotificationStatuses []int8 `json:"notification_statuses" binding:"required"`
}

type GetTelegramGroupListRequest struct {
	CompanyID string `json:"company_id" form:"company_id" binding:"required"`
}
