package postgres

import (
	"bw-erp/models"
	"bw-erp/storage/repo"

	"github.com/jmoiron/sqlx"
)

type telegramGroupRepo struct {
	db *sqlx.DB
}

func NewTelegramGroupRepo(db *sqlx.DB) repo.TelegramGroupI {
	return &telegramGroupRepo{db: db}
}

func (stg *telegramGroupRepo) Create(entity models.CreateTelegramGroupRequest) error {
	_, err := stg.db.Exec(`INSERT INTO telegram_groups(
		chat_id,
		code,
		name
	) VALUES (
		$1,
		$2,
		$3
	)`,
		entity.ChatID,
		entity.Code,
		entity.Name,
	)

	if err != nil {
		return err
	}
	return nil
}

func (stg *telegramGroupRepo) GetNotificationGroup(CompanyID string, Status int) (models.TelegramGroup, error) {
	var group models.TelegramGroup
	query := `select chat_id,with_location from telegram_groups where company_id = $1 and $2 = any(notification_statuses) limit 1 `

	err := stg.db.QueryRow(query, CompanyID, Status).Scan(
		&group.ChatID,
		&group.WithLocation,
	)
	if err != nil {
		return group, err
	}
	return group, nil
}
