package postgres

import (
	"bw-erp/models"
	"bw-erp/storage/repo"
	"encoding/json"
	"strings"

	"github.com/jmoiron/sqlx"
)

type statusChangeHistoryRepo struct {
	db *sqlx.DB
}

func NewStatusChangeHistoryRepo(db *sqlx.DB) repo.StatusChangeHistoryI {
	return &statusChangeHistoryRepo{db: db}
}

func (stg *statusChangeHistoryRepo) Create(entity models.CreateStatusChangeHistoryModel) (int, error) {
	stg.db.Exec(`INSERT INTO status_change_histories(
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
		entity.HistoryDetails.Status,
	)

	var (
		user_ids        string
		notification_id int
	)

	stg.db.QueryRow(`select user_ids from notification_settings where company_id = $1 and status = $2`, entity.CompanyID, entity.HistoryDetails.Status).Scan(&user_ids)

	if user_ids != "" {
		user_ids = strings.Trim(user_ids, "{}")
		userIDArray := strings.Split(user_ids, ",")

		DetailsJson, err := json.Marshal(entity.HistoryDetails)

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
			entity.HistoryableID,
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

	return notification_id, nil
}
