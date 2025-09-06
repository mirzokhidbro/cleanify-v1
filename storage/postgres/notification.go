package postgres

import (
	"bw-erp/models"
	"bw-erp/storage/repo"
	"encoding/json"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type notificationRepo struct {
	db *sqlx.DB
}

func NewNotificationRepo(db *sqlx.DB) repo.NotificationI {
	return &notificationRepo{db: db}
}

func (stg notificationRepo) Create(entity models.CreateNotificationModel) (notification_id int, err error) {
	var (
		user_ids string
	)

	if entity.Details.Type == "status_changed" {
		stg.db.QueryRow(`select user_ids from notification_settings where company_id = $1 and status = $2`, entity.CompanyID, entity.Details.Status).Scan(&user_ids)

		if user_ids != "" {
			user_ids = strings.Trim(user_ids, "{}")
			userIDArray := strings.Split(user_ids, ",")

			DetailsJson, _ := json.Marshal(entity.Details)

			err = stg.db.QueryRow(`INSERT INTO notifications(
				company_id,
				model_type,
				model_id,
				details
			) VALUES (
				$1,
				$2,
				$3,
				$4
			) RETURNING id`,
				entity.CompanyID,
				"orders",
				entity.ModelID,
				DetailsJson,
			).Scan(&notification_id)

			if err != nil {
				return notification_id, err
			}

			for _, user_id := range userIDArray {

				_, err = stg.db.Exec(`INSERT INTO user_notifications(
					notification_id,
					user_id
				) VALUES (
					$1,
					$2
				)`,
					&notification_id,
					&user_id,
				)

				if err != nil {
					return notification_id, err
				}
			}
		}
	}

	if entity.Details.Type == "order_attached" {
		DetailsJson, _ := json.Marshal(entity.Details)

		err = stg.db.QueryRow(`INSERT INTO notifications(
			company_id,
			model_type,
			model_id,
			details
		) VALUES (
			$1,
			$2,
			$3,
			$4
		) RETURNING id`,
			entity.CompanyID,
			"orders",
			entity.ModelID,
			DetailsJson,
		).Scan(&notification_id)

		if err != nil {
			return notification_id, err
		}

		_, err = stg.db.Exec(`INSERT INTO user_notifications(
			notification_id,
			user_id
		) VALUES (
			$1,
			$2
		)`,
			&notification_id,
			&entity.Details.Courier,
		)

		if err != nil {
			return notification_id, err
		}

	}

	return notification_id, nil
}

func (stg notificationRepo) GetMyNotifications(entity models.GetMyNotificationsRequest) ([]models.GetMyNotificationsResponse, error) {
	rows, err := stg.db.Query(`SELECT ui.id, n.company_id, n.model_type, n.model_id, n.details, ui.created_at 
								FROM user_notifications ui
								INNER JOIN notifications n ON ui.notification_id = n.id
								WHERE ui.user_id = $1 AND n.company_id = $2 and ui.created_at::date = now()::date order by ui.created_at desc`, entity.UserID, entity.CompanyID)

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

func (stg notificationRepo) GetNotificationsByID(entity models.GetNotificationsByIDRequest) ([]models.GetMyNotificationsResponse, error) {

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

func (stg notificationRepo) GetUnreadNotificationsCount(userID int64) (int, error) {
	var count int
	err := stg.db.QueryRow(`
		SELECT COUNT(*) 
		FROM user_notifications 
		WHERE user_id = $1 AND is_read = false and created_at::date = now()::date`, userID).Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}
