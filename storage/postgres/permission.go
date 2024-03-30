package postgres

import (
	"bw-erp/helper"
	"bw-erp/models"
	"bw-erp/pkg/utils"
	"bw-erp/storage/repo"
	"errors"

	"github.com/jmoiron/sqlx"
)

type permissionRepo struct {
	db *sqlx.DB
}

func NewPermissionRepo(db *sqlx.DB) repo.PermissionI {
	return &permissionRepo{db: db}
}

func (stg *permissionRepo) GetList(Scope string) ([]models.Permission, error) {
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
func (stg *permissionRepo) GetByPrimaryKey(ID string) (models.Permission, error) {
	var permission models.Permission
	if !utils.IsValidUUID(ID) {
		return permission, errors.New("permission id si noto'g'ri")
	}
	err := stg.db.QueryRow(`select id, slug, name from permissions where id = $1`, ID).Scan(
		&permission.ID,
		&permission.Slug,
		&permission.Name,
	)
	if err != nil {
		return permission, errors.New("permission topilmadi")
	}
	return permission, nil
}
