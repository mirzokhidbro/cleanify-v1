package models

type CreateCompanyRoleModel struct {
	Name      string `json:"name" binding:"required" minLength:"2" maxLength:"255"`
	CompanyId string `json:"company_id" binding:"required"`
}

type CompanyRoleListByCompany struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CompanyID string `json:"company_id"`
}
