package storage

import (
	"bw-erp/models"
)

type StorageI interface {
	CreateUserModel(id string, entity models.CreateUserModel) error
}
