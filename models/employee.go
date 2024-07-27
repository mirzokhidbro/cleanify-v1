package models

type Employee struct {
	CompanyID string `json:"company_id"`
	Phone     string `json:"phone"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type CreateEmployeeRequest struct {
	CompanyID string `json:"company_id" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
}

type ShowEmployeeRequest struct {
	CompanyID  string `json:"company_id" form:"company_id" binding:"required"`
	EmployeeID int    `json:"id" form:"id" binding:"required"`
}

type GetEmployeeList struct {
	CompanyID string `json:"company_id"`
	Phone     string `json:"phone"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	ID        int    `json:"id"`
}

type GetEmployeeListRequest struct {
	CompanyID string `json:"company_id" form:"company_id" binding:"required"`
}

type ShowEmployeeResponse struct {
	ID        int    `json:"id"`
	CompanyID string `json:"company_id"`
	Phone     string `json:"phone"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type EmployeeTransactionRequest struct {
	Salary        float64 `json:"salary"`
	ReceivedMoney float64 `json:"received_money" binding:"required"`
	EmployeeID    int     `json:"employee_id" binding:"required"`
	CompanyID     string  `json:"company_id" binding:"required"`
	UserID        string
}
