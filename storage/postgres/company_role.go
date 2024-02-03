package postgres

import (
	"bw-erp/models"
)

func (stg *Postgres) CreateCompanyRoleModel(id string, entity models.CreateRoleModel) error {
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
