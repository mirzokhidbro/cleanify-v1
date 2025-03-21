package postgres

import (
	"bw-erp/models"
	"bw-erp/pkg/utils"
	"bw-erp/storage/repo"
	"errors"

	"github.com/jmoiron/sqlx"
)

type roleRepo struct {
	db *sqlx.DB
}

func NewRoleRepo(db *sqlx.DB) repo.RoleI {
	return &roleRepo{db: db}
}

func (stg *roleRepo) Create(id string, entity models.CreateRoleModel) error {
	// _, _ = stg.db.Exec(`INSERT INTO roles(
	// 	id,
	// 	name,
	// 	company_id
	// ) VALUES (
	// 	$1,
	// 	$2,
	// 	$3
	// )`,
	// 	id,
	// 	entity.Name,
	// 	entity.CompanyId,
	// )

	// query := `DELETE FROM "role_and_permissions" WHERE role_id = $1`

	// _, err := stg.db.Exec(query, id)
	// if err != nil {
	// 	return err
	// }

	// PermissionIDs := utils.SetArray(entity.PermissionIDs)
	// _, err = stg.db.Exec(`INSERT INTO role_and_permissions(
	// 	role_id,
	// 	permission_ids
	// ) VALUES (
	// 	$1,
	// 	$2
	// )`,
	// 	id,
	// 	PermissionIDs,
	// )

	// if err != nil {
	// 	return err
	// }
	return nil
}

func (stg *roleRepo) GetListByCompany(companyID string) ([]models.RoleListByCompany, error) {
	rows, err := stg.db.Query(`select id, name, company_id from roles where company_id = $1`, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []models.RoleListByCompany
	for rows.Next() {
		var role models.RoleListByCompany
		err = rows.Scan(&role.ID, &role.Name, &role.CompanyID)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func (stg *roleRepo) GetByPrimaryKey(roleID string) (models.RoleByPrimaryKey, error) {
	// var model models.GetRoleByPrimaryKey
	var response models.RoleByPrimaryKey
	// err := stg.db.QueryRow(`select r.id, r.name, rp.permission_ids from roles r left join role_and_permissions rp on r.id::text = rp.role_id where r.id::text = $1`, roleID).Scan(
	// 	&model.ID,
	// 	&model.Name,
	// 	&model.PermissionIDs,
	// )
	// if err != nil {
	// 	return response, err
	// }
	// if *model.PermissionIDs != "" {
	// 	permissionIds := utils.GetArray(*model.PermissionIDs)
	// 	response.PermissionIDs = &permissionIds
	// }
	// response.ID = model.ID
	// response.Name = model.Name

	return response, nil
}

func (stg *roleRepo) GetPermissionsToRole(entity models.GetPermissionToRoleRequest) error {
	// for _, permission_id := range entity.PermissionIDs {
	// 	_, err := stg.GetPermissionByPrimaryKey(permission_id)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	// query := `DELETE FROM "role_and_permissions" WHERE role_id = $1`

	// _, err := stg.db.Exec(query, entity.RoleID)
	// if err != nil {
	// 	return err
	// }
	// PermissionIDs := utils.SetArray(entity.PermissionIDs)
	// _, err = stg.db.Exec(`INSERT INTO role_and_permissions(
	// 	role_id,
	// 	permission_ids
	// ) VALUES (
	// 	$1,
	// 	$2
	// )`,
	// 	entity.RoleID,
	// 	PermissionIDs,
	// )

	// if err != nil {
	// 	return err
	// }
	return nil
}

func (stg *roleRepo) GetPermissionByPrimaryKey(ID string) (models.Permission, error) {
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
