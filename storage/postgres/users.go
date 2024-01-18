package postgres

import (
	"bw-erp/models"
	"bw-erp/utils"
)

func (stg Postgres) CreateUserModel(id string, entity models.CreateUserModel) error {
	password, _ := utils.HashPassword(entity.Password)
	_, err := stg.db.Exec(`INSERT INTO users(
		id,
		phone,
		firstname,
		lastname,
		password
	) VALUES (
		$1,
		$2,
		$3, 
		$4,
		$5
	)`,
		id,
		entity.Phone,
		entity.Firstname,
		entity.Lastname,
		password,
	)

	if err != nil {
		return err
	}

	return err
}
