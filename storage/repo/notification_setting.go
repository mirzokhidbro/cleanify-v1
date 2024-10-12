package repo

import "bw-erp/models"

type NotificationSettingI interface {
	NotificationSetting(entity models.SetNotificationSettingRequest) error
	UsersListForNotificationSettings(companyID string) []models.UsersListForNotificationSettings
	GetUsersByStatus(entity models.GetUsersByStatusRequest) (models.GetUsersByStatus, error)
}
