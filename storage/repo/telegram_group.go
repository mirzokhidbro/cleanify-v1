package repo

import "bw-erp/models"

type TelegramGroupI interface {
	Create(entity models.CreateTelegramGroupRequest) error
	GetNotificationGroup(CompanyID string, Status int) (models.TelegramGroup, error)
}
