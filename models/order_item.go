package models

type OrderItemStatus int8

const (
	OrderItemStatusDraft    OrderItemStatus = 0
	OrderItemStatusWashed   OrderItemStatus = 1
	OrderItemStatusPrepared OrderItemStatus = 2
)

type CreateOrderItemModel struct {
	OrderID         int     `json:"order_id" binding:"required"`
	OrderItemTypeID string  `json:"order_item_type_id" binding:"required"`
	Price           float32 `json:"price" binding:"required"`
	Width           float32 `json:"width"`
	Height          float32 `json:"height"`
	OrderItemStatus int8    `json:"status"`
	Description     string  `json:"description"`
	ItemType        string  `json:"item_type"`
	IsCountable     bool    `json:"is_countable"`
}

type UpdateOrderItemRequest struct {
	ID              int     `json:"id" binding:"required"`
	OrderItemTypeID string  `json:"order_item_type_id" binding:"required"`
	Price           float32 `json:"price"`
	Width           float32 `json:"width"`
	Height          float32 `json:"height"`
	Description     string  `json:"description"`
	ItemType        string  `json:"item_type"`
	IsCountable     bool    `json:"is_countable"`
}

type UpdateOrderItemStatusRequest struct {
	OrderItemStatus OrderItemStatus `json:"status" binding:"required"`
	ID              int             `json:"id" binding:"required"`
}

type OrderItem struct {
	ID                  int                   `json:"id"`
	OrderID             int                   `json:"order_id"`
	Type                string                `json:"type"`
	Price               float32               `json:"price"`
	Width               float32               `json:"width"`
	Height              float32               `json:"height"`
	OrderItemStatus     OrderItemStatus       `json:"status"`
	IsCountable         bool                  `json:"is_countable"`
	Description         string                `json:"description"`
	OrderItemTypeID     string                `json:"order_item_type_id"`
	StatusChangeHistory []StatusChangeHistory `json:"status_change_history"`
}
