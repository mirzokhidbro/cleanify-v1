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

func (stg Postgres) GetUserByPhone(phone string) (models.AuthUserModel, error) {
	var user models.AuthUserModel
	err := stg.db.QueryRow(`SELECT id, phone, password from users where phone = $1`, phone).Scan(
		&user.ID,
		&user.Phone,
		&user.Password,
	)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (stg Postgres) GetUserById(id string) (models.User, error) {
	var user models.User
	err := stg.db.QueryRow(`select id, firstname, lastname, phone from users where id = $1`, id).Scan(
		&user.ID,
		&user.Firstname,
		&user.Lastname,
		&user.Phone,
	)
	if err != nil {
		return user, err
	}

	return user, nil
}
