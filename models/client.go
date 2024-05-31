package models

import (
	"time"
)

type CreateClientModel struct {
	CompanyID             string  `json:"company_id" binding:"required"`
	Address               string  `json:"address" binding:"required" minLength:"2" maxLength:"255"`
	FullName              string  `json:"full_name"`
	PhoneNumber           string  `json:"phone_number" binding:"required"`
	AdditionalPhoneNumber string  `json:"additional_phone_number"`
	WorkNumber            string  `json:"work_number"`
	Latitute              float64 `json:"latitute"`
	Longitude             float64 `json:"longitude"`
}

type ClientListRequest struct {
	Limit     int32  `json:"limit" form:"limit"`
	Offset    int32  `json:"offset" form:"offset"`
	Phone     string `json:"status,omitempty" form:"phone"`
	Address   string `json:"slug,omitempty" form:"address"`
	CompanyID string `json:"company_id" form:"company_id" binding:"required"`
}

type ClientList struct {
	ID                    int       `json:"id"`
	Address               string    `json:"address"`
	FullName              string    `json:"full_name"`
	PhoneNumber           string    `json:"phone_number"`
	AdditionalPhoneNumber string    `json:"additional_phone_number"`
	WorkNumber            string    `json:"work_number"`
	CreatedAt             time.Time `json:"created_at"`
}

type ClientListResponse struct {
	Data  []ClientList `json:"data"`
	Count int          `json:"total"`
}

type Client struct {
	ID                    int      `json:"id"`
	Address               string   `json:"address" binding:"required" minLength:"2" maxLength:"255"`
	FullName              string   `json:"full_name"`
	PhoneNumber           string   `json:"phone_number" binding:"required"`
	AdditionalPhoneNumber string   `json:"additional_phone_number"`
	WorkNumber            string   `json:"work_number"`
	Latitute              *float64 `json:"latitute"`
	Longitude             *float64 `json:"longitude"`
}

type GetClientByPrimaryKeyResponse struct {
	Client
	Orders []OrderLink `json:"orders"`
}

type OrderLink struct {
	ID        int       `json:"id"`
	Count     int       `json:"count"`
	Slug      string    `json:"slug"`
	Square    float64   `json:"square"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateClientRequest struct {
	ID                    int     `json:"id" binding:"required"`
	CompanyID             string  `json:"company_id" binding:"required"`
	FullName              string  `json:"full_name"`
	PhoneNumber           string  `json:"phone_number"`
	AdditionalPhoneNumber string  `json:"additional_phone_number"`
	WorkNumber            string  `json:"work_number"`
	Address               string  `json:"address"`
	Latitute              float64 `json:"latitute"`
	Longitude             float64 `json:"longitude"`
}
