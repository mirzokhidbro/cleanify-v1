package postgres

import (
	"bw-erp/models"
	"bw-erp/storage/repo"
	"strings"

	"github.com/jmoiron/sqlx"
)

type statusChangeHistoryRepo struct {
	db *sqlx.DB
}

func NewStatusChangeHistoryRepo(db *sqlx.DB) repo.StatusChangeHistoryI {
	return &statusChangeHistoryRepo{db: db}
}

func (stg *statusChangeHistoryRepo) Create(entity models.CreateStatusChangeHistoryModel) error {
	_, err := stg.db.Exec(`INSERT INTO status_change_histories(
		historyable_id,
		historyable_type,
		user_id,
		status
	) VALUES (
		$1,
		$2,
		$3,
		$4
	)`,
		entity.HistoryableID,
		entity.HistoryableType,
		entity.UserID,
		entity.Status,
	)

	var (
		user_ids        string
		notification_id int
	)

	stg.db.QueryRow(`select user_ids from notification_settings where company_id = $1 and status = $2`, entity.CompanyID, entity.Status).Scan(&user_ids)

	if user_ids != "" {
		user_ids = strings.Trim(user_ids, "{}")
		userIDArray := strings.Split(user_ids, ",")

		err := stg.db.QueryRow(`INSERT INTO notifications(
			company_id,
			model_type,
			model_id,
			status
		) VALUES (
			$1,
			$2,
			$3,
			$4
		) RETURNING id`,
			entity.CompanyID,
			"orders",
			entity.HistoryableID,
			entity.Status,
		).Scan(&notification_id)

		if err != nil {
			return err
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
				return err
			}
		}

	}

	return err
}
