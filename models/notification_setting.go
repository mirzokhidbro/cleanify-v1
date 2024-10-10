package models

type SetNotificationSettingRequest struct {
	UserID    string `json:"user_id" binding:"required"`
	CompanyID string `json:"company_id" binding:"required"`
	Statuses  []int8 `json:"statuses"`
}
