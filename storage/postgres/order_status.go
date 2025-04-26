package postgres

import (
	"bw-erp/helper"
	"bw-erp/models"
	"bw-erp/storage/repo"

	"github.com/jmoiron/sqlx"
)

type orderStatusRepo struct {
	db *sqlx.DB
}

func NewOrderStatusRepo(db *sqlx.DB) repo.OrderStatusI {
	return &orderStatusRepo{db: db}
}

func (stg *orderStatusRepo) GetList(companyID string, sortBy string, sortOrder string) (res []models.OrderStatus, err error) {
	var orderStatuses []models.OrderStatus
	var arr []interface{}
	params := make(map[string]interface{})

	query := "select id, number, name, coalesce(color, ''), description, slug, \"order\" from order_statuses"
	filter := " WHERE true"
	order := " ORDER BY "

	if sortBy != "" && (sortBy == "number" || sortBy == "order") {
		if sortBy == "order" {
			sortBy = "\"order\""
		}
		if sortOrder != "" && (sortOrder == "asc" || sortOrder == "desc") {
			order += sortBy + " " + sortOrder
		} else {
			order += sortBy + " asc"
		}
	} else {
		order += "\"order\" asc"
	}

	params["company_id"] = companyID
	filter += " AND (company_id = :company_id)"

	q := query + filter + order

	q, arr = helper.ReplaceQueryParams(q, params)
	rows, err := stg.db.Query(q, arr...)
	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		var orderStatus models.OrderStatus
		err = rows.Scan(
			&orderStatus.ID,
			&orderStatus.Number,
			&orderStatus.Name,
			&orderStatus.Color,
			&orderStatus.Description,
			&orderStatus.Slug,
			&orderStatus.Order,
		)
		if err != nil {
			return res, err
		}
		orderStatuses = append(orderStatuses, orderStatus)
	}

	if err = rows.Err(); err != nil {
		return res, err
	}

	return orderStatuses, nil
}

func (stg *orderStatusRepo) Reorder(companyID string, orders []models.OrderStatusOrder) error {
	tx, err := stg.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	for _, order := range orders {
		_, err = tx.Exec(
			`UPDATE order_statuses SET "order" = $1, updated_at = now() WHERE id = $2 AND company_id = $3`,
			order.Order,
			order.ID,
			companyID,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (stg orderStatusRepo) Update(entity models.UpdateOrderStatusRequest) (rowsAffected int64, err error) {
	query := `UPDATE "order_statuses" SET `

	if entity.Color != "" {
		query += `color = :color,`
	}
	if entity.Name != "" {
		query += `name = :name,`
	}

	query += `updated_at = now()
			  WHERE	id = :id and company_id = :company_id`

	params := map[string]interface{}{
		"id":         entity.ID,
		"name":       entity.Name,
		"color":      entity.Color,
		"company_id": entity.CompanyID,
	}

	query, arr := helper.ReplaceQueryParams(query, params)
	result, err := stg.db.Exec(query, arr...)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (stg orderStatusRepo) GetById(id int) (models.OrderStatus, error) {
	var order_status models.OrderStatus
	err := stg.db.QueryRow(`select id, number, name, coalesce(color, ''), description from order_statuses where id = $1`, id).Scan(
		&order_status.ID,
		&order_status.Number,
		&order_status.Name,
		&order_status.Color,
		&order_status.Description,
	)
	if err != nil {
		return order_status, err
	}

	return order_status, nil
}
