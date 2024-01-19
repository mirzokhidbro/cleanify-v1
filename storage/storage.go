package storage

import (
	"bw-erp/models"
)

type StorageI interface {
	CreateUserModel(id string, entity models.CreateUserModel) error
	GetUserByPhone(phone string) (models.AuthUserModel, error)
	GetUserById(id string) (models.User, error)
	GetUsersList() ([]models.User, error)
}
