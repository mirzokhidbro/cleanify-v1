package models

type GetMyNotificationsRequest struct {
	CompanyID string `json:"company_id" form:"company_id" binding:"required"`
	UserID    string `json:"user_id" form:"user_id" binding:"required"`
}

type GetMyNotificationsResponse struct {
	CompanyID string `json:"company_id"`
	ModelType string `json:"model_type"`
	Status    int    `json:"status"`
}
