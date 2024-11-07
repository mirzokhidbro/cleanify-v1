package postgres

import (
	"bw-erp/models"
	"bw-erp/storage/repo"

	"github.com/jmoiron/sqlx"
)

type notificationRepo struct {
	db *sqlx.DB
}

func NewNotificationRepo(db *sqlx.DB) repo.NotificationI {
	return &notificationRepo{db: db}
}

func (stg notificationRepo) GetMyNotifications(entity models.GetMyNotificationsRequest) ([]models.GetMyNotificationsResponse, error) {

	rows, err := stg.db.Query(`select n.company_id, n.model_type, n.status from user_notifications ui
								inner join notifications n on ui.notification_id = n.id
								where ui.user_id = $1 and n.company_id = $2`, entity.UserID, entity.CompanyID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []models.GetMyNotificationsResponse
	for rows.Next() {
		var notification models.GetMyNotificationsResponse
		err = rows.Scan(
			&notification.CompanyID,
			&notification.ModelType,
			&notification.Status)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return notifications, nil
}

func (stg notificationRepo) GetMyLatestNotifications(entity models.GetMyNotificationsRequest) (models.GetMyNotificationsResponse, error) {
	var res models.GetMyNotificationsResponse
	err := stg.db.QueryRow(`select n.company_id, n.model_type, n.status from user_notifications ui
								inner join notifications n on ui.notification_id = n.id
								where ui.user_id = $1 and n.company_id = $2`, entity.UserID, entity.CompanyID).Scan(&res.CompanyID, &res.ModelType, &res.Status)

	if err != nil {
		return res, err
	}

	return res, err

}

func (stg notificationRepo) GetNotificationsByStatus(entity models.GetNotificationsByStatusRequest) ([]models.GetMyNotificationsResponse, error) {

	rows, err := stg.db.Query(`select n.company_id, n.model_type, n.status, ui.user_id from user_notifications ui
								inner join notifications n on ui.notification_id = n.id
								where n.id = $1 `, entity.NotificationID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []models.GetMyNotificationsResponse
	for rows.Next() {
		var notification models.GetMyNotificationsResponse
		err = rows.Scan(
			&notification.CompanyID,
			&notification.ModelType,
			&notification.Status,
			&notification.UserID)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return notifications, nil
}
