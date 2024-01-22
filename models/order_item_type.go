package models

type OrderItemTypeModel struct {
	Name      string  `json:"name" binding:"required" minLength:"2" maxLength:"255"`
	Price     float32 `json:"price" binding:"required"`
	CopmanyID string  `json:"company_id" binding:"required"`
}

type OrderItemByCompany struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	CopmanyName string  `json:"company_name"`
	CopmanyID   string  `json:"company_id"`
}
