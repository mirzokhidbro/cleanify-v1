package repo

import "bw-erp/models"

type TelegramGroupI interface {
	Create(entity models.CreateTelegramGroupRequest) error
	GetNotificationGroup(CompanyID string, Status int) (models.TelegramGroup, error)
	Verification(Code int, companyID string) (models.TelegramGroup, error)
	GetList(companyId string) ([]models.TelegramGroupGetListResponse, error)
	GetByPrimaryKey(id int) (models.TelegramGroupGetByPrimayKeyResponse, error)
}
