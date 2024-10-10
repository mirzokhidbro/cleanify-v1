package postgres

import (
	"bw-erp/models"
	"bw-erp/pkg/utils"
	"bw-erp/storage/repo"

	"github.com/jmoiron/sqlx"
)

type notificationSettingRepo struct {
	db *sqlx.DB
}

func NewNotificationSettingRepo(db *sqlx.DB) repo.NotificationSettingI {
	return &notificationSettingRepo{db: db}
}

func (repo *notificationSettingRepo) NotificationSetting(entity models.SetNotificationSettingRequest) error {
	notification_statuses := utils.SetArray(utils.IntSliceToInterface(entity.Statuses))
	_, err := repo.db.Exec(`
		INSERT INTO notification_settings (user_id, company_id, statuses)	
		VALUES ($1, $2, $3)
		ON CONFLICT (user_id, company_id)
		DO UPDATE SET
		user_id = EXCLUDED.user_id,
		company_id = EXCLUDED.company_id,
		statuses = EXCLUDED.statuses`,
		entity.UserID,
		entity.CompanyID,
		notification_statuses,
	)
	if err != nil {
		return err
	}
	return nil
}
