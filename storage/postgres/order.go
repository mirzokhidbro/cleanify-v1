package postgres

import (
	"bw-erp/helper"
	"bw-erp/models"
	"errors"
)

func (stg *Postgres) CreateOrderModel(entity models.CreateOrderModel) (id int, err error) {
	_, err = stg.GetCompanyById(entity.CompanyID)
	if err != nil {
		return 0, errors.New("company not found")
	}

	err = stg.db.QueryRow(`INSERT INTO orders(
		company_id,
		phone,
		count,
		slug,
		description,
		address,
		client_id
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7
	) RETURNING id`,
		entity.CompanyID,
		entity.Phone,
		entity.Count,
		entity.Slug,
		entity.Description,
		entity.Address,
		entity.ClientID,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (stg *Postgres) GetOrdersList(companyID string, queryParam models.OrdersListRequest) (res models.OrderListResponse, err error) {
	var arr []interface{}
	res = models.OrderListResponse{}
	params := make(map[string]interface{})
	query := `SELECT 
		o.id, 
		o.slug, 
		o.status, 
		o.address,
		o.created_at,
		coalesce(sum(oi.price*oi.width*oi.height), 0) as price, 
		coalesce(sum(oi.width*oi.height), 0) as square 
		FROM "orders" as o left join order_items oi on o.id = oi.order_id`

	filter := " WHERE true"
	group := " group by o.id, o.slug, o.status, o.address, o.created_at"
	order := " ORDER BY created_at"
	arrangement := " DESC"
	offset := " OFFSET 0"
	limit := " LIMIT 20"

	params["company_id"] = companyID
	filter += " and (company_id = :company_id)"

	// if len(queryParam.ID) > 0 {
	// 	params["id"] = queryParam.ID
	// 	filter += " AND (('id') ILIKE ('%' || :id || '%'))"
	// }

	if queryParam.ID != 0 {
		params["id"] = queryParam.ID
		filter += " AND (o.id = :id)"
	}

	if queryParam.Status != 0 {
		params["status"] = queryParam.Status
		filter += " AND (o.status = :status)"
	}

	if !queryParam.DateFrom.IsZero() {
		params["date_from"] = queryParam.DateFrom
		filter += " AND (o.created_at >= :date_from::date)"
	}

	if !queryParam.DateTo.IsZero() {
		params["date_to"] = queryParam.DateTo
		filter += " AND (o.created_at <= :date_to::date)"
	}

	if queryParam.Offset > 0 {
		params["offset"] = queryParam.Offset
		offset = " OFFSET :offset"
	}

	if queryParam.Limit > 0 {
		params["limit"] = queryParam.Limit
		limit = " LIMIT :limit"
	}
	cQ := `SELECT count(1) FROM "orders" as o` + filter
	cQ, arr = helper.ReplaceQueryParams(cQ, params)
	err = stg.db.QueryRow(cQ, arr...).Scan(
		&res.Count,
	)

	if err != nil {
		return res, err
	}

	q := query + filter + group + order + arrangement + offset + limit

	q, arr = helper.ReplaceQueryParams(q, params)
	rows, err := stg.db.Query(q, arr...)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.OrderList
		err = rows.Scan(
			&order.ID,
			&order.Slug,
			&order.Status,
			&order.Address,
			&order.CreatedAt,
			&order.Price,
			&order.Square,
		)
		if err != nil {
			return res, err
		}
		res.Data = append(res.Data, order)
	}

	if err = rows.Err(); err != nil {
		return res, err
	}

	return res, nil
}

func (stg *Postgres) GetOrdersByStatus(companyID string, Status int) (order []models.Order, err error) {
	rows, err := stg.db.Query(`select  
		id,
		address, 
		phone
	from orders 
	where company_id = $1 and status = $2`, companyID, Status)
	if err != nil {
		return nil, err
	}

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		err = rows.Scan(
			&order.ID,
			&order.Address,
			&order.PhoneNumber,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (stg *Postgres) GetOrderByPhone(companyID string, Phone string) (models.Order, error) {
	var order models.Order
	err := stg.db.QueryRow(`select id, company_id, phone, count, slug, description, latitute, longitude, created_at, updated_at from orders where company_id = $1 and phone = $2`, companyID, Phone).Scan(
		&order.ID,
		&order.CompanyID,
		&order.PhoneNumber,
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

	return order, nil
}

func (stg *Postgres) GetOrderDetailedByPrimaryKey(ID int) (models.OrderShowResponse, error) {
	var order models.OrderShowResponse
	err := stg.db.QueryRow(`select o.id, 
									o.company_id, 
									COALESCE(c.phone_number, ''), 
									COALESCE(c.additional_phone_number, ''), 
									COALESCE(c.work_number, ''), 
									o.count, 
									o.slug, 
									o.description, 
									c.latitute, 
									c.longitude, 
									COALESCE(o.client_id, 0), 
									COALESCE(o.address, ''),
									o.status,
									o.created_at,
									o.updated_at 
								from orders o
								left join clients c on o.client_id = c.id 
								where o.id = $1`, ID).Scan(
		&order.ID,
		&order.CompanyID,
		&order.PhoneNumber,
		&order.AdditionalPhoneNumber,
		&order.WorkNumber,
		&order.Count,
		&order.Slug,
		&order.Description,
		&order.Latitute,
		&order.Longitude,
		&order.ClientID,
		&order.Address,
		&order.Status,
		&order.CreatedAt,
		&order.UpdatedAt,
	)
	if err != nil {
		return order, err
	}

	rows, err := stg.db.Query(`select id, order_id, type, price, width, height, description from order_items where order_id = $1`, ID)
	if err != nil {
		return order, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.OrderItem
		if err := rows.Scan(&item.ID, &item.OrderID, &item.Type, &item.Price, &item.Width, &item.Height, &item.Description); err != nil {
			return order, err
		}
		order.OrderItems = append(order.OrderItems, item)
	}

	return order, nil
}

func (stg *Postgres) GetOrderByPrimaryKey(ID int) (models.OrderShowResponse, error) {
	var order models.OrderShowResponse
	err := stg.db.QueryRow(`select o.id, 
									o.company_id, 
									COALESCE(c.phone_number, ''), 
									COALESCE(c.additional_phone_number, ''), 
									COALESCE(c.work_number, ''), 
									o.count, 
									o.slug, 
									o.description, 
									c.latitute, 
									c.longitude, 
									COALESCE(o.client_id, 0), 
									COALESCE(o.address, ''),
									o.created_at,
									o.updated_at 
								from orders o
								left join clients c on o.client_id = c.id 
								where o.id = $1`, ID).Scan(
		&order.ID,
		&order.CompanyID,
		&order.PhoneNumber,
		&order.AdditionalPhoneNumber,
		&order.WorkNumber,
		&order.Count,
		&order.Slug,
		&order.Description,
		&order.Latitute,
		&order.Longitude,
		&order.ClientID,
		&order.Address,
		&order.CreatedAt,
		&order.UpdatedAt,
	)
	if err != nil {
		return order, err
	}

	return order, nil
}

func (stg *Postgres) UpdateOrder(entity *models.UpdateOrderRequest) (rowsAffected int64, err error) {
	query := `UPDATE "orders" SET `

	if entity.Slug != "" {
		query += `slug = :slug,`
	}
	if entity.Status != 0 {
		query += `status = :status,`
	}
	if entity.Phone != "" {
		query += `phone = :phone,`
	}
	if entity.Count != 0 {
		query += `count = :count,`
	}
	if entity.Description != "" {
		query += `description = :description,`
	}

	query += `updated_at = now()
			  WHERE
					id = :id`

	order, _ := stg.GetOrderByPrimaryKey(entity.ID)
	if entity.Longitude != 0 && entity.Latitute != 0 && order.ClientID != 0 {
		updateOrderQuery := `UPDATE "clients" SET longitude = :longitude, latitute = :latitute WHERE id = :clientId`
		clientParams := map[string]interface{}{
			"clientId":  order.ClientID,
			"longitude": entity.Longitude,
			"latitute":  entity.Latitute,
		}
		updateOrderQuery, arr := helper.ReplaceQueryParams(updateOrderQuery, clientParams)
		_, err := stg.db.Exec(updateOrderQuery, arr...)
		if err != nil {
			return 0, err
		}

	}

	params := map[string]interface{}{
		"id":          entity.ID,
		"status":      entity.Status,
		"slug":        entity.Slug,
		"phone":       entity.Phone,
		"description": entity.Description,
		"count":       entity.Count,
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

func (stg *Postgres) GetOrderLocation(ID int) (models.Order, error) {
	var order models.Order
	err := stg.db.QueryRow(`select o.id, 
								   o.company_id, 
								   c.phone_number, 
								   c.additional_phone_number, 
								   c.work_number, 
								   o.count, 
								   o.slug, 
								   o.description, 
								   c.latitute, 
								   c.longitude, 
								   o.created_at, 
								   o.updated_at 
								from orders o
							left join clients c on o.client_id = c.id 
							where o.id = $1`, ID).Scan(
		&order.ID,
		&order.CompanyID,
		&order.PhoneNumber,
		&order.AdditionalPhoneNumber,
		&order.WorkNumber,
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

	return order, nil
}
