package postgres

import (
	"bw-erp/models"
	"bw-erp/utils"
	"errors"
)

func (stg Postgres) CreateUserModel(id string, entity models.CreateUserModel) error {
	if entity.ConfirmationPassword != entity.Password {
		return errors.New("confirmation password is not the same with password!")
	}
	password, _ := utils.HashPassword(entity.Password)
	_, err := stg.db.Exec(`INSERT INTO users(
		id,
		phone,
		firstname,
		lastname,
		role_id,
		password
	) VALUES (
		$1,
		$2,
		$3, 
		$4,
		$5,
		$6
	)`,
		id,
		entity.Phone,
		entity.Firstname,
		entity.Lastname,
		entity.RoleID,
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
	err := stg.db.QueryRow(`select u.id, u.firstname, u.lastname, u.phone, c.name, cr.name, c.id  from users u left join company_roles cr on cr.id = u.role_id left join companies c on c.id = cr.company_id where u.id = $1`, id).Scan(
		&user.ID,
		&user.Firstname,
		&user.Lastname,
		&user.Phone,
		&user.Company,
		&user.Role,
		&user.CompanyID,
	)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (stg Postgres) GetUsersList() ([]models.User, error) {
	rows, err := stg.db.Query(`select 
									u.id, 
									u.firstname, 
									u.lastname, 
									u.phone, 
									c.name, 
									cr.name  
									from users u 
									left join company_roles cr on cr.id = u.role_id 
									left join companies c on c.id = cr.company_id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(
			&user.ID,
			&user.Firstname,
			&user.Lastname,
			&user.Phone,
			&user.Company,
			&user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
