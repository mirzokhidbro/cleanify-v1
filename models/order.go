package models

type CreateOrderModel struct {
	CompanyID   string `json:"company_id", binding:"required"`
	Phone       string `json:"phone", binding:"required"`
	Count       int    `json:"count", binding:"required"`
	Slug        string `json:"slug", binding:"required"`
	Description string `json:"description"`
}

type OrderList struct {
	ID     int        `json:"id"`
	Slug   string     `json:"slug"`
	Status NullString `json:"status"`
}
