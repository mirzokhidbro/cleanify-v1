package postgres

import (
	"bw-erp/models"
)

func (stg *Postgres) CreateCompanyRoleModel(id string, entity models.CreateCompanyRoleModel) error {
	_, err := stg.GetCompanyById(entity.CompanyId)
	if err != nil {
		return err
	}

	_, err = stg.db.Exec(`INSERT INTO company_roles(
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

func (stg *Postgres) GetRolesListByCompany(companyID string) ([]models.CompanyRoleListByCompany, error) {
	rows, err := stg.db.Query(`select id, name, company_id from company_roles where company_id = $1`, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companyRoles []models.CompanyRoleListByCompany
	for rows.Next() {
		var companyRole models.CompanyRoleListByCompany
		err = rows.Scan(&companyRole.ID, &companyRole.Name, &companyRole.CompanyID)
		if err != nil {
			return nil, err
		}
		companyRoles = append(companyRoles, companyRole)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return companyRoles, nil
}
