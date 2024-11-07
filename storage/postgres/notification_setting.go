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
	user_ids := utils.SetArray(utils.StringSliceToInterface(entity.UserIDs))
	_, err := repo.db.Exec(`
		INSERT INTO notification_settings (user_ids, company_id, status)	
		VALUES ($1, $2, $3)
		ON CONFLICT (status, company_id)
		DO UPDATE SET
		user_ids = EXCLUDED.user_ids,
		company_id = EXCLUDED.company_id,
		status = EXCLUDED.status`,
		user_ids,
		entity.CompanyID,
		entity.Status,
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

func (repo *notificationSettingRepo) GetUsersByStatus(entity models.GetUsersByStatusRequest) (models.GetUsersByStatus, error) {
	var usersByStatus models.GetUsersByStatus

	err := repo.db.QueryRow(`select name, number from order_statuses where company_id = $1 and number = $2`, entity.CompanyID, entity.Status).Scan(&usersByStatus.StatusName, &usersByStatus.StatusNumber)

	if err != nil {
		return usersByStatus, err
	}

	rows, err := repo.db.Query(`SELECT u.firstname || ' ' || lastname as fullname, id
								FROM users u
								JOIN (
									SELECT DISTINCT unnest(ns.user_ids) AS user_id
									FROM order_statuses os
									JOIN notification_settings ns ON os.company_id = ns.company_id AND ns.status = os.number
									WHERE os.number = $1 AND os.company_id = $2
								) AS subquery ON u.id = subquery.user_id::uuid`, entity.Status, entity.CompanyID)

	if err != nil {
		return usersByStatus, err
	}

	for rows.Next() {
		var user models.UsersListForNotificationSettings
		if err := rows.Scan(&user.Fullname, &user.UserID); err != nil {
			return usersByStatus, err
		}

		usersByStatus.Users = append(usersByStatus.Users, user)
	}

	return usersByStatus, nil

}
