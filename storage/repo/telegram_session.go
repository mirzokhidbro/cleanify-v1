package repo

import "bw-erp/models"

type TelegramSessionI interface {
	GetByChatIDBotID(ChatID int64, BotID int64) (models.TelegramSessionModel, error)
	Delete(ID int) (rowsAffected int64, err error)
	Create(entity models.TelegramSessionModel) error
}
