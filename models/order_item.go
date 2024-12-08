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
	ID                  int                   `json:"id,omitempty"`
	OrderID             int                   `json:"order_id,omitempty"`
	Type                string                `json:"type,omitempty"`
	Price               float32               `json:"price,omitempty"`
	Width               float32               `json:"width,omitempty"`
	Height              float32               `json:"height,omitempty"`
	OrderItemStatus     OrderItemStatus       `json:"status,omitempty"`
	IsCountable         bool                  `json:"is_countable,omitempty"`
	Description         string                `json:"description,omitempty"`
	OrderItemTypeID     string                `json:"order_item_type_id,omitempty"`
	StatusChangeHistory []StatusChangeHistory `json:"status_change_history,omitempty"`
}
