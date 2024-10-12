package models

type SetNotificationSettingRequest struct {
	UserIDs   []string `json:"user_ids" binding:"required"`
	CompanyID string   `json:"company_id" binding:"required"`
	Status    int8     `json:"status"`
}

type UsersListForNotificationSettings struct {
	Fullname string `json:"fullname"`
	UserID   string `json:"user_id"`
}

type UsersListForNotificationSettingsRequest struct {
	CompanyID string `json:"company_id" form:"company_id" binding:"required"`
}

type GetUsersByStatus struct {
	StatusNumber int8                               `json:"status_number"`
	StatusName   string                             `json:"status_name"`
	Users        []UsersListForNotificationSettings `json:"users"`
}

type GetUsersByStatusRequest struct {
	Status    int8   `json:"status" form:"status" binding:"required"`
	CompanyID string `json:"company_id" form:"company_id" binding:"required"`
}
