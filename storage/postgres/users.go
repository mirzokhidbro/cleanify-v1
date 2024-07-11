package postgres

import (
	"bw-erp/helper"
	"bw-erp/models"
	"bw-erp/pkg/utils"
	"bw-erp/storage/repo"
	"errors"
	"strings"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) repo.UserI {
	return &userRepo{db: db}
}

func (stg userRepo) Create(id string, entity models.CreateUserModel) error {
	if entity.ConfirmationPassword != entity.Password {
		return errors.New("confirmation password is not the same with password")
	}
	password, _ := utils.HashPassword(entity.Password)
	PermissionIDs := utils.SetArray(utils.StringSliceToInterface(entity.PermissionIDs))

	_, err := stg.db.Exec(`INSERT INTO users(
		id,
		phone,
		firstname,
		lastname,
		password,
		permission_ids,
		company_id
	) VALUES (
		$1,
		$2,
		$3, 
		$4,
		$5,
		$6,
		$7
	)`,
		id,
		entity.Phone,
		entity.Firstname,
		entity.Lastname,
		password,
		PermissionIDs,
		entity.CompanyID,
	)

	if err != nil {
		return err
	}

	return err
}

func (stg userRepo) GetByPhone(phone string) (models.AuthUserModel, error) {
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

func (stg userRepo) GetById(id string) (models.User, error) {
	var user models.User
	err := stg.db.QueryRow(`select u.id, u.firstname, u.lastname, u.phone, c.id from users u left join companies c on c.id = u.company_id where u.id = $1`, id).Scan(
		&user.ID,
		&user.Firstname,
		&user.Lastname,
		&user.Phone,
		&user.CompanyID,
	)
	if err != nil {
		return user, err
	}

	rows, err := stg.db.Query(`select c.id, c.name, up.permission_ids from user_permissions up 
									inner join companies c on c.id = up.company_id 
									where up.user_id = $1 order by c.name desc`, user.ID)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	for rows.Next() {
		var permissions models.UserPermissionByCompany
		if err := rows.Scan(&permissions.CompanyID, &permissions.CompanyName, &permissions.Can); err != nil {
			return user, err
		}

		if permissions.Can != "" {
			Permissions := utils.GetArray(permissions.Can)
			Permission := ""
			for _, permissionID := range Permissions {
				permission, err := stg.GetPermissionByPrimaryKey(permissionID.(string))
				if err == nil {
					Permission += "|" + permission.Slug
				}
			}
			permissions.Can = strings.TrimPrefix(Permission, "|")
			permissions.PermissionIDs = utils.InterfaceSliceToString(Permissions)
		}

		user.UserPermissionByCompany = append(user.UserPermissionByCompany, permissions)
	}

	return user, nil
}

func (stg userRepo) GetList(companyID string) ([]models.User, error) {
	rows, err := stg.db.Query(`SELECT 
								u.id, 
								u.firstname, 
								u.lastname, 
								u.phone,
								c.id
								FROM users u 
								LEFT JOIN companies c ON c.id = u.company_id 
								WHERE c.id is not null and c.id = $1`, companyID)

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
			&user.CompanyID)
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

func (stg userRepo) ChangePassword(userID string, entity models.ChangePasswordRequest) error {
	if entity.NewPassword != entity.NewPasswordConfirmation {
		return errors.New("confirmation password is not the same with password")
	}
	password, _ := utils.HashPassword(entity.NewPassword)

	query := `UPDATE "users" SET
	password = :password,
	updated_at = now()
	WHERE
	id = :id`

	params := map[string]interface{}{
		"id":       userID,
		"password": password,
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	_, err := stg.db.Exec(q, arr...)
	if err != nil {
		return err
	}
	return err
}

func (stg *userRepo) GetPermissionByPrimaryKey(ID string) (models.Permission, error) {
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

func (stg *userRepo) Edit(entity models.UpdateUserRequest) (rowsAffected int64, err error) {
	query := `UPDATE "users" SET `

	if entity.Firstname != "" {
		query += `firstname = :firstname,`
	}
	if entity.Lastname != "" {
		query += `lastname = :lastname,`
	}

	query += `updated_at = now()
			  WHERE id = :id `

	params := map[string]interface{}{
		"id":        entity.ID,
		"firstname": entity.Firstname,
		"lastname":  entity.Lastname,
	}

	query, arr := helper.ReplaceQueryParams(query, params)

	result, err := stg.db.Exec(query, arr...)

	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	// [TODO: refactoring]

	for _, permission := range entity.Permissions {
		var UserPermissionByCompany models.UserPermissionByCompany
		stg.db.QueryRow(`select company_id from user_permissions where company_id = $1 and user_id = $2`, permission.CompanyID, entity.ID).Scan(
			&UserPermissionByCompany.CompanyID,
		)

		PermissionIDs := utils.SetArray(utils.StringSliceToInterface(permission.PermissionIDs))

		if len(UserPermissionByCompany.CompanyID) > 0 {
			query = `UPDATE "user_permissions" SET permission_ids = :permission_ids, updated_at = now() where company_id = :company_id and user_id = :user_id`

			permissionEditParams := map[string]interface{}{
				"permission_ids": PermissionIDs,
				"company_id":     permission.CompanyID,
				"user_id":        entity.ID,
			}

			query, arr := helper.ReplaceQueryParams(query, permissionEditParams)

			_, err := stg.db.Exec(query, arr...)

			if err != nil {
				return 0, err
			}
		} else {
			_, err := stg.db.Exec(`INSERT INTO user_permissions(
				permission_ids,
				company_id,
				user_id
			) VALUES (
				$1,
				$2,
				$3
			)`,
				PermissionIDs,
				permission.CompanyID,
				entity.ID,
			)

			if err != nil {
				return 0, err
			}
		}
	}

	return rowsAffected, nil
}
