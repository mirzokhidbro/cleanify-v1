package postgres

import (
	"bw-erp/models"
	"bw-erp/storage/repo"

	"github.com/jmoiron/sqlx"
)

type statusChangeHistoryRepo struct {
	db *sqlx.DB
}

func NewStatusChangeHistoryRepo(db *sqlx.DB) repo.StatusChangeHistoryI {
	return &statusChangeHistoryRepo{db: db}
}

func (stg *statusChangeHistoryRepo) Create(entity models.CreateStatusChangeHistoryModel) error {
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
		entity.Status,
	)

	return nil
}
