package postgres

import (
	"bw-erp/helper"
	"bw-erp/models"
	"bw-erp/storage/repo"
	"errors"

	"github.com/jmoiron/sqlx"
)

type clientRepo struct {
	db *sqlx.DB
}

func NewClientRepo(db *sqlx.DB) repo.ClientStorageI {
	return &clientRepo{db: db}
}

func (stg *clientRepo) Create(entity models.CreateClientModel) (id int, err error) {
	err = stg.db.QueryRow(`INSERT INTO clients(
		company_id,
		address,
		full_name,
		phone_number,
		additional_phone_number,
		work_number,
		longitude,
		latitute
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8
	) RETURNING id`,
		entity.CompanyID,
		entity.Address,
		entity.FullName,
		entity.PhoneNumber,
		entity.AdditionalPhoneNumber,
		entity.WorkNumber,
		entity.Longitude,
		entity.Latitute,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (stg *clientRepo) GetList(companyID string, queryParam models.ClientListRequest) (res models.ClientListResponse, err error) {
	var arr []interface{}
	res = models.ClientListResponse{}
	params := make(map[string]interface{})
	query := `SELECT 
		id, 
		address, 
		COALESCE(full_name, ''), 
		phone_number,
		COALESCE(additional_phone_number, ''),
		COALESCE(work_number, ''),
		created_at 
		FROM "clients"`

	filter := " WHERE true"
	order := " ORDER BY created_at"
	arrangement := " DESC"
	offset := " OFFSET 0"
	limit := " LIMIT 20"

	params["company_id"] = companyID
	filter += " and (company_id = :company_id)"

	if len(queryParam.Phone) > 3 {
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

func (stg *clientRepo) GetByPrimaryKey(ID int) (models.GetClientByPrimaryKeyResponse, error) {
	var client models.GetClientByPrimaryKeyResponse
	err := stg.db.QueryRow(`select id, address, COALESCE(full_name, ''), phone_number, COALESCE(additional_phone_number, ''), COALESCE(work_number, ''), latitute, longitude from clients where id = $1`, ID).Scan(
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

	rows, err := stg.db.Query(`select 
									o.id, 
									o.count, 
									coalesce(o.slug, ''), 
									o.created_at,
									ROUND(CAST(COALESCE(sum(oi.price*oi.width*oi.height), 0) AS NUMERIC), 2) as price, 
									round(cast(coalesce(sum(oi.width*oi.height), 0) as numeric), 2) as square 
								from orders as o 
								left join order_items oi on o.id = oi.order_id
								where client_id = $1 group by o.id, o.count, o.slug, o.created_at`, ID)
	if err != nil {
		return client, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.OrderLink
		if err := rows.Scan(&item.ID, &item.Count, &item.Slug, &item.CreatedAt, &item.Price, &item.Square); err != nil {
			return client, err
		}
		client.Orders = append(client.Orders, item)
	}

	return client, nil
}

func (stg *clientRepo) Update(companyID string, entity models.UpdateClientRequest) (rowsAffected int64, err error) {
	query := `UPDATE "clients" SET `

	if entity.Longitude != 0 {
		query += `longitude = :longitude,`
	}
	if entity.Latitute != 0 {
		query += `latitute = :latitute,`
	}
	if entity.FullName != "" {
		query += `full_name = :full_name,`
	}
	if entity.PhoneNumber != "" {
		query += `phone_number = :phone_number,`
	}
	if entity.AdditionalPhoneNumber != "" {
		query += `additional_phone_number = :additional_phone_number,`
	}
	if entity.WorkNumber != "" {
		query += `work_number = :work_number,`
	}
	if entity.Address != "" {
		query += `address = :address,`
	}

	query += `updated_at = now()
			  WHERE
					id = :id and company_id = :company_id`

	params := map[string]interface{}{
		"id":                      entity.ID,
		"longitude":               entity.Longitude,
		"latitute":                entity.Latitute,
		"full_name":               entity.FullName,
		"phone_number":            entity.PhoneNumber,
		"additional_phone_number": entity.AdditionalPhoneNumber,
		"work_number":             entity.WorkNumber,
		"company_id":              companyID,
		"address":                 entity.Address,
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

func (stg *clientRepo) GetByPhoneNumber(Phone string, CompanyID string) (models.Client, error) {
	var client models.Client
	err := stg.db.QueryRow(`select id, address, COALESCE(full_name, ''), phone_number, COALESCE(additional_phone_number, ''), COALESCE(work_number, ''), latitute, longitude from clients where phone_number = $1 and company_id = $2`, Phone, CompanyID).Scan(
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

	return client, nil
}
