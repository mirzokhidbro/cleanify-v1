package models

type CreateRoleModel struct {
	Name          string   `json:"name" binding:"required" minLength:"2" maxLength:"255"`
	CompanyId     string   `json:"company_id" binding:"required"`
	PermissionIDs []string `json:"permission_ids" binding:"required"`
}

type RoleListByCompany struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CompanyID string `json:"company_id"`
}

type GetPermissionToRoleRequest struct {
	CompanyID     string   `json:"company_id"`
	RoleID        string   `json:"role_id" binding:"required"`
	PermissionIDs []string `json:"permission_ids" binding:"required"`
}

// [TODO: refactoring!]
type GetRoleByPrimaryKey struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	PermissionIDs string `json:"permission_ids"`
}

type RoleByPrimaryKey struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	PermissionIDs []string `json:"permission_ids"`
}
