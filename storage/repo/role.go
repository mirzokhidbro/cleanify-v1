package repo

import "bw-erp/models"

type RoleI interface {
	Create(id string, entity models.CreateRoleModel) error
	GetListByCompany(companyID string) ([]models.RoleListByCompany, error)
	GetPermissionsToRole(models.GetPermissionToRoleRequest) error
	GetByPrimaryKey(roleID string) (models.RoleByPrimaryKey, error)
}
