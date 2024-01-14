package model

import "time"

type BotUser struct {
	ID          int       `gorm:"primaryKey"`
	ChatID      uint64    `gorm:"uniqueIndex;not null"`
	PhoneNumber string    `gorm:"type:varchar(16)"`
	FirstName   string    `gorm:"type:varchar(255);"`
	LastName    string    `gorm:"type:varchar(255);"`
	Username    string    `gorm:"type:varchar(255);"`
	Role        string    `gorm:"string"`
	Page        string    `gorm:"type:varchar(255)"`
	DialogStep  string    `gorm:"type:varchar(255)"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
}
