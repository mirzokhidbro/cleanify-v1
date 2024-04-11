package models

type CreateOrderItemModel struct {
	OrderID         int     `json:"order_id" binding:"required"`
	OrderItemTypeID string  `json:"order_item_type_id" binding:"required"`
	Price           float32 `json:"price" binding:"required"`
	Width           float32 `json:"width"`
	Height          float32 `json:"height"`
	Description     string  `json:"description"`
	ItemType        string  `json:"item_type"`
	IsCountable     bool    `json:"is_countable"`
	Count           int     `json:"count"`
}

type UpdateOrderItemRequest struct {
	ID          int     `json:"id" binding:"required"`
	Price       float32 `json:"price"`
	Width       float32 `json:"width"`
	Height      float32 `json:"height"`
	Description string  `json:"description"`
	Type        string  `json:"types"`
}

type OrderItem struct {
	ID          int     `json:"id"`
	OrderID     int     `json:"order_id"`
	Type        string  `json:"type"`
	Price       float32 `json:"price"`
	Width       float32 `json:"width"`
	Height      float32 `json:"height"`
	IsCountable bool    `json:"is_countable"`
	Count       int     `json:"count"`
	Description string  `json:"description"`
}
