package repo

import "bw-erp/models"

type TelegramBotI interface {
	Create(CompanyID string, entity models.CreateCompanyBotModel) error
	GetByCompany(CompanyID string) (models.CompanyTelegramBot, error)
	GetOrderBot() ([]models.CompanyTelegramBot, error)
}
