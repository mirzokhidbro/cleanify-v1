package models

type OrderStatus struct {
	ID          int    `json:"id"`
	Number      int    `json:"number"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
}

type UpdateOrderStatusRequest struct {
	ID        int    `json:"id"`
	CompanyID string `json:"company_id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
}

type GetOrderStatusListRequest struct {
	CompanyID string `json:"company_id" form:"company_id" binding:"required"`
}
