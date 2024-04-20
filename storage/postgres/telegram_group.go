package postgres

import (
	"bw-erp/helper"
	"bw-erp/models"
	"bw-erp/storage/repo"
	"errors"

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

func (stg *telegramGroupRepo) Verification(Code int, companyID string) (models.TelegramGroup, error) {
	var group models.TelegramGroup
	query := `select chat_id, code from telegram_groups where code = $1 and company_id is null`

	err := stg.db.QueryRow(query, Code).Scan(
		&group.ChatID,
		&group.Code,
	)
	if err != nil {
		return group, err
	}

	if group.Code == 0 {
		return group, errors.New("guruh topilmadi")
	}

	query = `UPDATE "telegram_groups" SET updated_at = now(), company_id = :company_id where code = :code`

	params := map[string]interface{}{
		"code":       Code,
		"company_id": companyID,
	}

	query, arr := helper.ReplaceQueryParams(query, params)
	result, err := stg.db.Exec(query, arr...)
	if err != nil {
		return group, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return group, err
	}

	return group, nil
}

func (stg *telegramGroupRepo) GetList(companyId string) ([]models.TelegramGroupGetListResponse, error) {
	var groups []models.TelegramGroupGetListResponse
	rows, err := stg.db.Query(`select id, company_id, name, notification_statuses, with_location, created_at, updated_at from telegram_groups where company_id = $1`, companyId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var group models.TelegramGroupGetListResponse
		err = rows.Scan(&group.ID, &group.CompanyID, &group.Name, &group.NotificationStatuses, &group.WithLocation, &group.CreatedAt, &group.UpdatedAt)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)

	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}
