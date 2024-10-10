package repo

import "bw-erp/models"

type NotificationSettingI interface {
	NotificationSetting(entity models.SetNotificationSettingRequest) error
	UsersListForNotificationSettings(companyID string) []models.UsersListForNotificationSettings
}
