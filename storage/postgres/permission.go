package postgres

import (
	"bw-erp/helper"
	"bw-erp/models"
	"bw-erp/utils"
	"errors"
)

func (stg *Postgres) GetPermissionList(Scope string) ([]models.Permission, error) {
	var arr []interface{}
	var permissions []models.Permission
	params := make(map[string]interface{})
	query := `SELECT 
				id,
				slug,
				name,
				scope
			 FROM "permissions"`

	filter := " WHERE true"

	if Scope != "super-admin" {
		params["scope"] = "company"
		filter += " AND (scope = :scope)"
	}

	q := query + filter

	q, arr = helper.ReplaceQueryParams(q, params)
	rows, err := stg.db.Query(q, arr...)
	if err != nil {
		return permissions, err
	}
	defer rows.Close()

	for rows.Next() {
		var permission models.Permission
		err = rows.Scan(
			&permission.ID,
			&permission.Slug,
			&permission.Name,
			&permission.Scope)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}

func (stg *Postgres) GetPermissionByPrimaryKey(ID string) (models.Permission, error) {
	var permission models.Permission
	if !utils.IsValidUUID(ID) {
		return permission, errors.New("Permission id si noto'g'ri!")
	}
	err := stg.db.QueryRow(`select id, slug, name from permissions where id = $1`, ID).Scan(
		&permission.ID,
		&permission.Slug,
		&permission.Name,
	)
	if err != nil {
		return permission, errors.New("Permission topilmadi!")
	}
	return permission, nil
}
