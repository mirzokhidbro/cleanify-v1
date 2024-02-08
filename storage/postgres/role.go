package postgres

import (
	"bw-erp/models"
	"bw-erp/utils"
)

func (stg *Postgres) CreateRoleModel(id string, entity models.CreateRoleModel) error {
	_, err := stg.GetCompanyById(entity.CompanyId)
	if err != nil {
		return err
	}

	_, err = stg.db.Exec(`INSERT INTO roles(
		id,
		name,
		company_id
	) VALUES (
		$1,
		$2,
		$3
	)`,
		id,
		entity.Name,
		entity.CompanyId,
	)

	query := `DELETE FROM "role_and_permissions" WHERE role_id = $1`

	_, err = stg.db.Exec(query, id)
	if err != nil {
		return err
	}

	PermissionIDs := utils.SetArray(entity.PermissionIDs)
	_, err = stg.db.Exec(`INSERT INTO role_and_permissions(
		role_id,
		permission_ids
	) VALUES (
		$1,
		$2
	)`,
		id,
		PermissionIDs,
	)

	if err != nil {
		return err
	}
	return nil
}

func (stg *Postgres) GetRolesListByCompany(companyID string) ([]models.RoleListByCompany, error) {
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

func (stg *Postgres) GetRoleByPrimaryKey(roleID string) (models.RoleByPrimaryKey, error) {
	var model models.GetRoleByPrimaryKey
	var response models.RoleByPrimaryKey
	err := stg.db.QueryRow(`select r.id, r.name, rp.permission_ids from roles r left join role_and_permissions rp on r.id::text = rp.role_id where r.id::text = $1`, roleID).Scan(
		&model.ID,
		&model.Name,
		&model.PermissionIDs,
	)
	if err != nil {
		return response, err
	}
	permissionIds := utils.GetArray(model.PermissionIDs)
	response.ID = model.ID
	response.Name = model.Name
	response.PermissionIDs = permissionIds


	return response, nil
}

func (stg *Postgres) GetPermissionsToRole(entity models.GetPermissionToRoleRequest) error {
	for _, permission_id := range entity.PermissionIDs {
		_, err := stg.GetPermissionByPrimaryKey(permission_id)
		if err != nil {
			return err
		}
	}
	query := `DELETE FROM "role_and_permissions" WHERE role_id = $1`

	_, err := stg.db.Exec(query, entity.RoleID)
	if err != nil {
		return err
	}
	PermissionIDs := utils.SetArray(entity.PermissionIDs)
	_, err = stg.db.Exec(`INSERT INTO role_and_permissions(
		role_id,
		permission_ids
	) VALUES (
		$1,
		$2
	)`,
		entity.RoleID,
		PermissionIDs,
	)

	if err != nil {
		return err
	}
	return nil
}
