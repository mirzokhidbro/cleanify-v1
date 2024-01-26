package postgres

import (
	"bw-erp/helper"
	"bw-erp/models"
	"errors"
	"math"
)

func (stg *Postgres) CreateOrderModel(entity models.CreateOrderModel) error {
	_, err := stg.GetCompanyById(entity.CompanyID)
	if err != nil {
		return errors.New("Company not found")
	}

	_, err = stg.db.Exec(`INSERT INTO orders(
		company_id,
		phone,
		count,
		slug,
		description
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5
	)`,
		entity.CompanyID,
		entity.Phone,
		entity.Count,
		entity.Slug,
		entity.Description,
	)

	if err != nil {
		return err
	}

	return nil
}

func (stg *Postgres) GetOrdersList(companyID string) ([]models.OrderList, error) {
	rows, err := stg.db.Query(`select id, slug, status from orders where company_id = $1`, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.OrderList
	for rows.Next() {
		var order models.OrderList
		err = rows.Scan(
			&order.ID,
			&order.Slug,
			&order.Status)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (stg *Postgres) GetOrderByPrimaryKey(ID int) (models.Order, error) {
	var order models.Order
	err := stg.db.QueryRow(`select id, company_id, phone, count, slug, description, latitute, longitude, created_at, updated_at from orders where id = $1`, ID).Scan(
		&order.ID,
		&order.CompanyID,
		&order.Phone,
		&order.Count,
		&order.Slug,
		&order.Description,
		&order.Latitute,
		&order.Longitude,
		&order.CreatedAt,
		&order.UpdatedAt,
	)
	if err != nil {
		return order, err
	}

	rows, err := stg.db.Query(`select order_id, type, price, width, height, description from order_items where order_id = $1`, ID)
	if err != nil {
		return order, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.OrderItem
		if err := rows.Scan(&item.OrderID, &item.Type, &item.Price, &item.Width, &item.Height, &item.Description); err != nil {
			return order, err
		}
		order.OrderItems = append(order.OrderItems, item)
	}

	return order, nil
}

func (stg *Postgres) UpdateOrder(entity *models.UpdateOrderRequest) (rowsAffected int64, err error) {
	query := `UPDATE "orders" SET `

	if entity.Slug != "" {
		query += `slug = :slug,`
	}
	if entity.Status == math.MaxInt16 {
		query += `status = :status,`
	}
	if entity.Phone != "" {
		query += `phone = :phone,`
	}
	if entity.Description != "" {
		query += `description = :description,`
	}

	query += `updated_at = now()
			  WHERE
					id = :id`

	params := map[string]interface{}{
		"id":          entity.ID,
		"status":      entity.Status,
		"slug":        entity.Slug,
		"phone":       entity.Phone,
		"description": entity.Description,
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
