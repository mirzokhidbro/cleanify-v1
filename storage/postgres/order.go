package postgres

import (
	"bw-erp/helper"
	"bw-erp/models"
	"errors"
	"fmt"
	"math"
)

func (stg *Postgres) CreateOrderModel(entity models.CreateOrderModel) (id int, err error) {
	_, err = stg.GetCompanyById(entity.CompanyID)
	fmt.Print("\n order create page")
	if err != nil {
		return 0, errors.New("Company not found")
	}

	err = stg.db.QueryRow(`INSERT INTO orders(
		company_id,
		phone,
		count,
		slug,
		description,
		chat_id
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6
	) RETURNING id`,
		entity.CompanyID,
		entity.Phone,
		entity.Count,
		entity.Slug,
		entity.Description,
		entity.ChatID,
	).Scan(&id)

	if err != nil {
		fmt.Print("\n order create error ", err)
		return 0, err
	}

	return id, nil
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
	if entity.Count != "" {
		query += `count = :count,`
	}
	if entity.Description != "" {
		query += `description = :description,`
	}
	if entity.Longitude != 0 {
		query += `longitude = :longitude,`
	}
	if entity.Latitute != 0 {
		query += `latitute = :latitute,`
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
		"count":       entity.Count,
		"longitude":   entity.Longitude,
		"latitute":    entity.Latitute,
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
