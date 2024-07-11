package repo

import "bw-erp/models"

type UserI interface {
	Create(id string, entity models.CreateUserModel) error
	GetByPhone(phone string) (models.AuthUserModel, error)
	GetById(id string) (models.User, error)
	GetList(companyID string) ([]models.User, error)
	ChangePassword(userID string, entity models.ChangePasswordRequest) error
	Edit(entity models.UpdateUserRequest) (rowsAffected int64, err error)
}
