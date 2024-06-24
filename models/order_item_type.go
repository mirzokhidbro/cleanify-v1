package models

type OrderItemTypeModel struct {
	Name        string  `json:"name" binding:"required,min=2,max=255"`
	Price       float32 `json:"price" binding:"required"`
	CompanyID   string  `json:"company_id" binding:"required"`
	IsCountable *bool   `json:"is_countable" binding:"required"`
}

type OrderItemByCompany struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	IsCountable bool    `json:"is_countable"`
	CopmanyName string  `json:"company_name"`
	CopmanyID   string  `json:"company_id"`
}

type EditOrderItemTypeRequest struct {
	CopmanyID   string  `json:"company_id" binding:"required"`
	ID          string  `json:"id" binding:"required"`
	Price       float32 `json:"price" binding:"required"`
	IsCountable *bool   `json:"is_countable" binding:"required"`
	Name        string  `json:"name" binding:"required"`
}

type GetOrderItemTypeByCompany struct {
	CompanyID string `json:"company_id" form:"company_id" binding:"required"`
}
