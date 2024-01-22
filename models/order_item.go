package models

type CreateOrderItemModel struct {
	OrderID     int     `json:"order_id", binding:"required"`
	Type        string  `json:"type", binding:"required"`
	Price       float32 `json:"price", binding:"required"`
	Width       float32 `json:"width", binding:"required"`
	Height      float32 `json:"height", binding:"required"`
	Description string  `json:"description"`
}

type OrderItem struct {
	OrderID     int     `json:"order_id"`
	Type        string  `json:"type"`
	Price       float32 `json:"price"`
	Width       float32 `json:"width"`
	Height      float32 `json:"height"`
	Description string  `json:"description"`
}
