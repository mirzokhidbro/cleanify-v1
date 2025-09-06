package repo

import "bw-erp/models"

type BotUserI interface {
	GetByChatID(ChatID int64, BotID int64) (models.BotUser, error)
	Create(entity models.CreateBotUserModel) error
	GetSelectedBotUser(BotID int64, Phone string) (models.SelectedUser, error)
	Update(entity models.BotUser) (rowsAffected int64, err error)
	GetByCompany(BotID int64, ChatID int64) (botUser models.BotUserByCompany, err error)
	GetByUserID(UserID int64) (models.BotUser, error)
	// GetNotificationGroup(CompanyID string) (models.BotUserByCompany, error)
}
