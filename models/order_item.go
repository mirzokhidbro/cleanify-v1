package models

type CreateOrderItemModel struct {
	OrderID         int     `json:"order_id", binding:"required"`
	OrderItemTypeID string  `json:"order_item_type_id", binding:"required"`
	Price           float32 `json:"price", binding:"required"`
	Width           float32 `json:"width", binding:"required"`
	Height          float32 `json:"height", binding:"required"`
	Description     string  `json:"description"`
}

type UpdateOrderItemRequest struct {
	ID          int     `json:"id", binding:"required"`
	Price       float32 `json:"price"`
	Width       float32 `json:"width"`
	Height      float32 `json:"height"`
	Description string  `json:"description"`
	Type        string  `json:"types"`
}

type OrderItem struct {
	OrderID     int     `json:"order_id"`
	Type        string  `json:"type"`
	Price       float32 `json:"price"`
	Width       float32 `json:"width"`
	Height      float32 `json:"height"`
	Description string  `json:"description"`
}
