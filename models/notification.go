package models

import "time"

type GetMyNotificationsRequest struct {
	CompanyID string `json:"company_id" form:"company_id" binding:"required"`
	UserID    string `json:"user_id" form:"user_id" binding:"required"`
}

type GetNotificationsByIDRequest struct {
	NotificationID int
	// Status    int8
	// CompanyID string
	// ModelType string
	// ModelID   int
}

type GetMyNotificationsResponse struct {
	CompanyID    string                   `json:"company_id"`
	ModelType    string                   `json:"model_type"`
	ModelID      int                      `json:"model_id"`
	UserID       string                   `json:"user_id"`
	PermformedAt time.Time                `json:"performed_at"`
	UnreadCount  int                      `json:"unread_count"`
	Details      OrderNotificationDetails `json:"details"`
}

type OrderNotificationDetails struct {
	Type    string `json:"type"`
	Address string `json:"address"`
	Status  int    `json:"status"`
}

type Message struct {
	ClientID string
	Content  string
}

type CreateNotificationModel struct {
	CompanyID string
	ModelType string
	ModelID   int
	Details   NotificationDetails
}

type NotificationDetails struct {
	Type    string `json:"type"`
	Address string `json:"address"`
	Status  int8   `json:"status"`
	Courier string `json:"courier"`
}
