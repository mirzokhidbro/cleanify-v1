package models

type OrderStatus struct {
	ID          int    `json:"id"`
	Number      int    `json:"number"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	Order       int16  `json:"order"`
}

type UpdateOrderStatusRequest struct {
	ID        int    `json:"id"`
	CompanyID string `json:"company_id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
}

type GetOrderStatusListRequest struct {
	CompanyID string `json:"company_id" form:"company_id" binding:"required"`
	SortBy    string `json:"sort_by" form:"sort_by" enums:"number,order"`
	SortOrder string `json:"sort_order" form:"sort_order" enums:"asc,desc"`
}

type OrderStatusOrder struct {
	ID    int   `json:"id" binding:"required"`
	Order int16 `json:"order" binding:"required"`
}

type ReorderOrderStatusRequest struct {
	CompanyID string             `json:"company_id" binding:"required"`
	Orders    []OrderStatusOrder `json:"orders" binding:"required,min=1"`
}
