package models

type GetMyNotificationsRequest struct {
	CompanyID string `json:"company_id" form:"company_id" binding:"required"`
	UserID    string `json:"user_id" form:"user_id" binding:"required"`
}

type GetNotificationsByStatusRequest struct {
	Status    int8
	CompanyID string
	ModelType string
	ModelID   int
}

type GetMyNotificationsResponse struct {
	CompanyID string `json:"company_id"`
	ModelType string `json:"model_type"`
	Status    int    `json:"status"`
	UserID    string `json:"user_id"`
}

type Message struct {
	ClientID string
	Content  string
}
