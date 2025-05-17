package postgres

import (
	"bw-erp/models"
	"bw-erp/storage/repo"

	"github.com/jmoiron/sqlx"
)

type companyRepo struct {
	db *sqlx.DB
}

func NewCompanyRepo(db *sqlx.DB) repo.CompanyStorageI {
	return &companyRepo{db: db}
}

func (stg *companyRepo) Create(id string, entity models.CreateCompanyModel) error {
	_, err := stg.db.Exec(`INSERT INTO companies(
		id,
		name,
		owner_id
	) VALUES (
		$1,
		$2,
		$3
	)`,
		id,
		entity.Name,
		entity.OwnerId,
	)

	if err != nil {
		return err
	}

	return nil
}

func (stg *companyRepo) GetById(id string) (models.Company, error) {
	var company models.Company
	err := stg.db.QueryRow(`select id, name from companies where id = $1`, id).Scan(
		&company.ID,
		&company.Name,
	)
	if err != nil {
		return company, err
	}

	return company, nil
}

func (stg *companyRepo) GetByOwnerId(ownerId string) ([]models.Company, error) {
	rows, err := stg.db.Query(`select id, name, owner_id from companies where owner_id = $1`, ownerId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companies []models.Company
	for rows.Next() {
		var company models.Company
		err = rows.Scan(&company.ID, &company.Name)
		if err != nil {
			return nil, err
		}
		companies = append(companies, company)

	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return companies, nil
}
