package postgres

import (
	"bw-erp/helper"
	"bw-erp/models"
)

func (stg *Postgres) GetPermissionList(Scope string) ([]models.Permissions, error) {
	var arr []interface{}
	var permissions []models.Permissions
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
		var permission models.Permissions
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
