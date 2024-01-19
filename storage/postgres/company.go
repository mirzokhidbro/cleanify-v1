package postgres

import "bw-erp/models"

func (stg Postgres) CreateCompanyModel(id string, entity models.CreateCompanyModel) error {

	_, err := stg.GetUserById(entity.OwnerId)
	if err != nil {
		return err
	}

	_, err = stg.db.Exec(`INSERT INTO companies(
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

	return err
}
