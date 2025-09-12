package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

type PaymentType int8
type PaymentStatus int8

const (
	Cach       PaymentType = 1 // cash
	CreditCard PaymentType = 2 // credit card
)

const (
	Pending PaymentStatus = 1
	Partial PaymentStatus = 2
	Paid    PaymentStatus = 3
)

type CreateOrderModel struct {
	CompanyID   string  `json:"company_id" binding:"required"`
	ClientID    int     `json:"client_id"`
	Phone       string  `json:"phone" binding:"required"`
	Count       int     `json:"count"`
	Slug        string  `json:"slug"`
	Status      int8    `json:"status"`
	Description string  `json:"description"`
	ChatID      int64   `json:"chat_id"`
	Address     string  `json:"address" binding:"required"`
	IsNewClient bool    `json:"is_new_client"`
	Latitute    float64 `json:"latitute"`
	Longitude   float64 `json:"longitude"`
}

type OrderList struct {
	ID          int       `json:"id"`
	OrderNumber *int      `json:"order_number"`
	Phone       string    `json:"phone"`
	Address     *string   `json:"address"`
	CourierID   *string   `json:"courier_id"`
	Status      int16     `json:"status"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
}

type OrderListResponse struct {
	Data  []OrderList `json:"data"`
	Count int         `json:"total"`
}

type OrdersListRequest struct {
	Limit         int32         `json:"limit" form:"limit"`
	Offset        int32         `json:"offset" form:"offset"`
	Status        int           `json:"status" form:"status"`
	PaymentStatus PaymentStatus `json:"payment_status" form:"payment_status"`
	Phone         string        `json:"phone" form:"phone"`
	Search        string        `json:"search" form:"search"`
	ID            string        `json:"id" form:"id"`
	DateFrom      time.Time     `json:"date_from" form:"date_from"`
	DateTo        time.Time     `json:"date_to" form:"date_to"`
	CompanyID     string        `json:"company_id" form:"company_id" binding:"required"`
	CourierID     string        `json:"courier_id" form:"courier_id"`
}

type OrderShowResponse struct {
	Order
	OrderItems       []OrderItem        `json:"order_items"`
	OrderTransaction []OrderTransaction `json:"transactions,omitempty"`
	Comments         []Comment          `json:"comment,omitempty"`
}

type CreateOrderComment struct {
	OrderID  int    `json:"order_id" form:"order_id" binding:"required"`
	Type     string `json:"type" form:"type" binding:"required,oneof=text voice"`
	Message  string `json:"message" form:"message"`
	VoiceURL string `json:"voice_url,omitempty"`
	UserID   int64  `json:"-"`
}

type OrderTransaction struct {
	ReceiverFullname string    `json:"receiver_fullname"`
	PaymentType      uint8     `json:"payment_type"`
	Amount           float64   `json:"amount"`
	CreatedAt        time.Time `json:"created_at"`
}

type Order struct {
	ID                    int                   `json:"id,omitempty"`
	CompanyID             string                `json:"company_id,omitempty"`
	Uuid                  string                `json:"uuid,omitempty"`
	ClientID              int                   `json:"client_id,omitempty"`
	CourierID             *int64                `json:"courier_id,omitempty"`
	PhoneNumber           string                `json:"phone_number,omitempty"`
	AdditionalPhoneNumber string                `json:"additional_phone_number,omitempty"`
	WorkNumber            string                `json:"work_number,omitempty"`
	Count                 int                   `json:"count,omitempty"`
	Slug                  string                `json:"slug,omitempty"`
	Status                int8                  `json:"status,omitempty"`
	Description           string                `json:"description,omitempty"`
	CreatedAt             time.Time             `json:"created_at,omitempty"`
	UpdatedAt             time.Time             `json:"updated_at,omitempty"`
	Latitute              *float64              `json:"latitute,omitempty"`
	Longitude             *float64              `json:"longitude,omitempty"`
	Address               *string               `json:"address,omitempty"`
	Square                float64               `json:"square,omitempty"`
	Price                 float64               `json:"price,omitempty"`
	StatusChangeHistory   []StatusChangeHistory `json:"status_change_history,omitempty"`
	PaymentStatus         int16                 `json:"payment_status,omitempty"`
	ServicePrice          float64               `json:"service_price,omitempty"`
	DiscountPercentage    float64               `json:"discount_percentage,omitempty"`
	DiscountPrice         float64               `json:"discounted_price,omitempty"`
}

type OrderReceipt struct {
	CompanyName        string      `json:"company_name,omitempty"`
	Address            string      `json:"address"`
	Phone              string      `json:"phone"`
	OrderNumber        int         `json:"order_number"`
	CreatedAt          time.Time   `json:"created_at"`
	ServicePrice       *float64    `json:"service_price"`
	DiscountPercentage *float64    `json:"discount_percentage"`
	DiscountedPrice    *float64    `json:"discounted_price"`
	OrderItems         []OrderItem `json:"order_item"`
}

type OrderSendLocationRequest struct {
	OrderID int `form:"order_id" binding:"required"`
}

type UpdateOrderRequest struct {
	ID              int           `json:"id" binding:"required"`
	CompanyID       string        `json:"company_id"`
	Address         string        `json:"address"`
	CourierID       int64         `json:"courier_id"`
	Slug            string        `json:"slug"`
	Status          int8          `json:"status"`
	PaymentStatus   PaymentStatus `json:"payment_status"`
	Phone           string        `json:"phone"`
	ChatID          int64         `json:"chat_id"`
	Description     string        `json:"description"`
	Count           int           `json:"count"`
	Latitute        float64       `json:"latitute"`
	Longitude       float64       `json:"longitude"`
	DiscountedPrice float64       `json:"discounted_price"`
}

type AddOrderPaymentRequest struct {
	CompanyID   string      `json:"company_id" binding:"required"`
	OrderID     int         `json:"order_id" binding:"required"`
	Amount      float64     `json:"amount" binding:"required"`
	PaymentType PaymentType `json:"payment_type" binding:"required,oneof=1 2"`
	Description string      `json:"description"`
}

type NullFloat struct {
	sql.NullFloat64
}

func (ns NullFloat) MarshalJSONFloat() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.Float64)
	}
	return json.Marshal(nil)
}

type DeleteOrderRequest struct {
	ID        int    `json:"id" binding:"required"`
	CompanyID string `json:"company_id" binding:"required"`
}

type SetOrderPriceRequest struct {
	ID        int    `json:"id" binding:"required"`
	CompanyID string `json:"company_id" binding:"required"`
	// ServicePrice float64 `json:"service_price" binding:"required"`
	// DiscountPercentage float64 `json:"discount_percentage" binding:"required"`
	DiscountedPrice float64 `json:"discounted_price"`
}
