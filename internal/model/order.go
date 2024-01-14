package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID        int `gorm:"primaryKey"`
	BotUserId uint `gorm:""`
	Count     int
	Latitude  float64
	Longitude float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
