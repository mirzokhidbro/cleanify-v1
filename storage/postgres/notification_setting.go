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

func (repo *notificationSettingRepo) UsersListForNotificationSettings(companyID string) []models.UsersListForNotificationSettings {
	var users []models.UsersListForNotificationSettings

	rows, err := repo.db.Query(`select firstname || ' ' || lastname as fullname, u.id from users u
												inner join user_permissions up on up.user_id = u.id
												where up.company_id = $1`, companyID)

	if err != nil {
		return users
	}
	defer rows.Close()

	for rows.Next() {
		var user models.UsersListForNotificationSettings
		if err := rows.Scan(&user.Fullname, &user.UserID); err != nil {
			return users
		}
		users = append(users, user)
	}

	return users
}
