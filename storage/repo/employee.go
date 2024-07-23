package repo

import "bw-erp/models"

type EmployeeI interface {
	Create(entity models.CreateEmployeeRequest) error
	GetList(companyID string) (res []models.GetEmployeeList, err error)
}
