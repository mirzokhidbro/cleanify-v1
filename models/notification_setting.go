package models

type SetNotificationSettingRequest struct {
	UserID    string `json:"user_id" binding:"required"`
	CompanyID string `json:"company_id" binding:"required"`
	Statuses  []int8 `json:"statuses"`
}

type UsersListForNotificationSettings struct {
	Fullname string `json:"fullname"`
	UserID   string `json:"user_id"`
}

type UsersListForNotificationSettingsRequest struct {
	CompanyID string `json:"company_id" form:"company_id" binding:"required"`
}
