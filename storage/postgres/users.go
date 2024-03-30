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
	err := stg.db.QueryRow(`select u.id, u.firstname, u.lastname, u.phone, c.name, r.name, c.id, r.id  from users u left join roles r on r.id = u.role_id left join companies c on c.id = r.company_id where u.id = $1`, id).Scan(
		&user.ID,
		&user.Firstname,
		&user.Lastname,
		&user.Phone,
		&user.Company,
		&user.Role,
		&user.CompanyID,
		&user.RoleID,
	)
	if err != nil {
		return user, err
	}

	if user.RoleID != nil {
		var roleAndPermission models.RoleAndPermission
		_ = stg.db.QueryRow(`select role_id, permission_ids from role_and_permissions where role_id = $1`, user.RoleID).Scan(
			&roleAndPermission.RoleID,
			&roleAndPermission.PermissionIDs,
		)
		if roleAndPermission.PermissionIDs != "" {
			permissionIds := utils.GetArray(roleAndPermission.PermissionIDs)
			Permission := ""
			for _, permissionID := range permissionIds {
				permission, err := stg.GetPermissionByPrimaryKey(permissionID)
				if err == nil {
					Permission += "|" + permission.Slug
				}
			}
			user.Permissions = strings.TrimPrefix(Permission, "|")
		}

	}

	return user, nil
}

func (stg userRepo) GetList(companyID string) ([]models.User, error) {
	rows, err := stg.db.Query(`SELECT 
								u.id, 
								u.firstname, 
								u.lastname, 
								u.phone, 
								c.name, 
								cr.name  
								FROM users u 
								LEFT JOIN roles cr ON cr.id = u.role_id 
								LEFT JOIN companies c ON c.id = cr.company_id 
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
