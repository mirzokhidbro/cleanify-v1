package postgres

import (
	"bw-erp/models"
	"bw-erp/storage/repo"
	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type notificationRepo struct {
	db *sqlx.DB
}

func NewNotificationRepo(db *sqlx.DB) repo.NotificationI {
	return &notificationRepo{db: db}
}

func (stg notificationRepo) GetMyNotifications(entity models.GetMyNotificationsRequest) ([]models.GetMyNotificationsResponse, error) {
	rows, err := stg.db.Query(`SELECT ui.id, n.company_id, n.model_type, n.model_id, n.details, ui.created_at 
								FROM user_notifications ui
								INNER JOIN notifications n ON ui.notification_id = n.id
								WHERE ui.user_id = $1 AND n.company_id = $2 and ui.created_at::date = now()::date`, entity.UserID, entity.CompanyID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []models.GetMyNotificationsResponse
	var notificationIDs []int
	for rows.Next() {
		var notification models.GetMyNotificationsResponse
		var detailsData []byte
		var notificationID int

		err = rows.Scan(
			&notificationID,
			&notification.CompanyID,
			&notification.ModelType,
			&notification.ModelID,
			&detailsData,
			&notification.PermformedAt)
		if err != nil {
			return nil, err
		}

		if len(detailsData) > 0 {
			err = json.Unmarshal(detailsData, &notification.Details)
			if err != nil {
				return nil, err
			}
		}

		notifications = append(notifications, notification)
		notificationIDs = append(notificationIDs, notificationID)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(notificationIDs) > 0 {
		query := `UPDATE user_notifications SET is_read = true WHERE id = ANY($1::int[])`
		_, err = stg.db.Exec(query, pq.Array(notificationIDs))
		if err != nil {
			return nil, err
		}
	}

	return notifications, nil
}

func (stg notificationRepo) GetMyLatestNotifications(entity models.GetMyNotificationsRequest) (models.GetMyNotificationsResponse, error) {
	var res models.GetMyNotificationsResponse
	err := stg.db.QueryRow(`select n.company_id, n.model_type, n.status from user_notifications ui
								inner join notifications n on ui.notification_id = n.id
								where ui.user_id = $1 and n.company_id = $2`, entity.UserID, entity.CompanyID).Scan(&res.CompanyID, &res.ModelType)

	if err != nil {
		return res, err
	}

	return res, err

}

func (stg notificationRepo) GetNotificationsByStatus(entity models.GetNotificationsByStatusRequest) ([]models.GetMyNotificationsResponse, error) {

	rows, err := stg.db.Query(`select ui.user_id, n.company_id, n.model_type, n.model_id, n.details, ui.created_at from user_notifications ui
								inner join notifications n on ui.notification_id = n.id
								where n.id = $1 `, entity.NotificationID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []models.GetMyNotificationsResponse
	for rows.Next() {
		var notification models.GetMyNotificationsResponse
		var details []byte
		err = rows.Scan(
			&notification.UserID,
			&notification.CompanyID,
			&notification.ModelType,
			&notification.ModelID,
			&details,
			&notification.PermformedAt)
		if err != nil {
			return nil, err
		}

		if len(details) > 0 {
			err = json.Unmarshal(details, &notification.Details)
			if err != nil {
				return nil, err
			}
		}

		notifications = append(notifications, notification)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return notifications, nil
}

func (stg notificationRepo) GetUnreadNotificationsCount(userID string, companyID string) (int, error) {
	var count int
	err := stg.db.QueryRow(`
		SELECT COUNT(*) 
		FROM user_notifications 
		WHERE user_id = $1 AND is_read = false and created_at::date = now()::date, company_id = $2`, userID, companyID).Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}
