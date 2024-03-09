package postgres

import (
	"bw-erp/helper"
	"bw-erp/models"
	"errors"
	"fmt"
)

func (stg *Postgres) CreateClientModel(entity models.CreateClientModel) (id int, err error) {
	_, err = stg.GetCompanyById(entity.CompanyID)
	if err != nil {
		return 0, errors.New("company not found")
	}

	err = stg.db.QueryRow(`INSERT INTO clients(
		company_id,
		address,
		full_name,
		phone_number,
		additional_phone_number,
		work_number
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6
	) RETURNING id`,
		entity.CompanyID,
		entity.Address,
		entity.FullName,
		entity.PhoneNumber,
		entity.AdditionalPhoneNumber,
		entity.WorkNumber,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (stg *Postgres) GetClientsList(companyID string, queryParam models.ClientListRequest) (res models.ClientListResponse, err error) {
	var arr []interface{}
	res = models.ClientListResponse{}
	params := make(map[string]interface{})
	query := `SELECT 
		id, 
		address, 
		full_name, 
		phone_number,
		additional_phone_number,
		work_number,
		created_at 
		FROM "clients"`

	filter := " WHERE true"
	order := " ORDER BY created_at"
	arrangement := " DESC"
	offset := " OFFSET 0"
	limit := " LIMIT 20"

	params["company_id"] = companyID
	filter += " and (company_id = :company_id)"

	if len(queryParam.Phone) > 0 {
		params["phone"] = queryParam.Phone
		filter += " AND ((phone_number || ' ' || additional_phone_number || ' ' || work_number) ILIKE ('%' || :phone || '%'))"
	}

	if len(queryParam.Address) > 0 {
		params["address"] = queryParam.Address
		filter += " AND ((address) ILIKE ('%' || :address || '%'))"
	}

	if queryParam.Offset > 0 {
		params["offset"] = queryParam.Offset
		offset = " OFFSET :offset"
	}

	if queryParam.Limit > 0 {
		params["limit"] = queryParam.Limit
		limit = " LIMIT :limit"
	}
	cQ := `SELECT count(1) FROM "clients"` + filter
	cQ, arr = helper.ReplaceQueryParams(cQ, params)
	err = stg.db.QueryRow(cQ, arr...).Scan(
		&res.Count,
	)

	if err != nil {
		return res, err
	}

	q := query + filter + order + arrangement + offset + limit

	q, arr = helper.ReplaceQueryParams(q, params)
	rows, err := stg.db.Query(q, arr...)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		var clients models.ClientList
		err = rows.Scan(
			&clients.ID,
			&clients.Address,
			&clients.FullName,
			&clients.PhoneNumber,
			&clients.AdditionalPhoneNumber,
			&clients.WorkNumber,
			&clients.CreatedAt)
		if err != nil {
			return res, err
		}
		res.Data = append(res.Data, clients)
	}

	if err = rows.Err(); err != nil {
		return res, err
	}

	return res, nil
}

func (stg *Postgres) GetClientByPrimaryKey(ID int) (models.GetClientByPrimaryKeyResponse, error) {
	var client models.GetClientByPrimaryKeyResponse
	fmt.Print("aa")
	err := stg.db.QueryRow(`select id, address, full_name, phone_number, additional_phone_number, work_number, latitute, longitude from clients where id = $1`, ID).Scan(
		&client.ID,
		&client.Address,
		&client.FullName,
		&client.PhoneNumber,
		&client.AdditionalPhoneNumber,
		&client.WorkNumber,
		&client.Latitute,
		&client.Longitude,
	)
	if err != nil {
		return client, errors.New("client not found")
	}

	rows, err := stg.db.Query(`select id, count, slug, created_at from orders where client_id = $1`, ID)
	if err != nil {
		return client, errors.New("error happened there 3")
	}
	defer rows.Close()

	for rows.Next() {
		var item models.OrderLink
		if err := rows.Scan(&item.ID, &item.Count, &item.Slug, &item.CreatedAt); err != nil {
			return client, errors.New("error happened there 3")
		}
		client.Orders = append(client.Orders, item)
	}

	return client, nil
}

func (stg *Postgres) UpdateClient(entity *models.UpdateClientRequest) (rowsAffected int64, err error) {
	query := `UPDATE "clients" SET `

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
		"id":        entity.ID,
		"longitude": entity.Longitude,
		"latitute":  entity.Latitute,
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
