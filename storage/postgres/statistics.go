package postgres

import (
	"bw-erp/helper"
	"bw-erp/models"
	"bw-erp/storage/repo"

	"github.com/jmoiron/sqlx"
)

type statisticsRepo struct {
	db *sqlx.DB
}

func NewStatisticsRepo(db *sqlx.DB) repo.StatisticsI {
	return &statisticsRepo{db: db}
}

func (stg *statisticsRepo) GetWorkVolume(companyID string) ([]models.WorkVolume, error) {
	var arr []interface{}
	var workVolumes []models.WorkVolume
	params := make(map[string]interface{})
	query := `SELECT
	round(sum((width::numeric * height::numeric)), 2) as meter_square,
		washed_at::date,
		type
		FROM order_items oi inner join orders o on oi.order_id = o.id`

	filter := " WHERE true"
	order := " ORDER BY washed_at"
	arrangement := " DESC"
	group := " group by washed_at::date, type"

	params["company_id"] = companyID
	filter += " AND (o.company_id = :company_id)"

	q := query + filter + group + order + arrangement

	q, arr = helper.ReplaceQueryParams(q, params)
	rows, err := stg.db.Query(q, arr...)
	if err != nil {
		return workVolumes, err
	}
	defer rows.Close()

	for rows.Next() {
		var workVolume models.WorkVolume
		err = rows.Scan(
			&workVolume.MeterSquare,
			&workVolume.WashedAt,
			&workVolume.Type)
		if err != nil {
			return nil, err
		}
		workVolumes = append(workVolumes, workVolume)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return workVolumes, nil
}
