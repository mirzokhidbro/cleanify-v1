package repo

import "bw-erp/models"

type PermissionI interface {
	GetList(Scope string) ([]models.Permission, error)
	GetByPrimaryKey(ID string) (models.Permission, error)
}
