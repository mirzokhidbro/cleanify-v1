package models

import (
	"encoding/json"
	"time"
)

type PushSubscription struct {
	ID        int64     `json:"id" db:"id"`
	UserID    string    `json:"user_id" db:"user_id"`
	Endpoint  string    `json:"endpoint" db:"endpoint"`
	AuthKey   string    `json:"auth_key" db:"auth_key"`
	P256dhKey string    `json:"p256dh_key" db:"p256dh_key"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type CreatePushSubscriptionRequest struct {
	UserID string          `json:"user_id" binding:"required"`
	Data   json.RawMessage `json:"data" binding:"required"`
}

type GetPushSubscriptionResponse struct {
	Subscriptions []PushSubscription `json:"subscriptions"`
	Count         int64              `json:"count"`
}
