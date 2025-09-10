package repo

import "bw-erp/models"

type UserI interface {
	Create(entity models.CreateUserModel) error
	GetByPhone(phone string) (models.AuthUserModel, error)
	GetById(id int64) (models.User, error)
	GetList(companyID string) ([]models.User, error)
	ChangePassword(userID int64, entity models.ChangePasswordRequest) error
	Edit(entity models.UpdateUserRequest) (rowsAffected int64, err error)
	GetCouriesList(companyID string) ([]models.GetCouriesResponse, error)
}
