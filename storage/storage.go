package storage

import (
	"bw-erp/models"
)

type StorageI interface {
	// user
	CreateUserModel(id string, entity models.CreateUserModel) error
	GetUserByPhone(phone string) (models.AuthUserModel, error)
	GetUserById(id string) (models.User, error)
	GetUsersList() ([]models.User, error)
	// company
	CreateCompanyModel(id string, entity models.CreateCompanyModel) error
	GetCompanyByOwnerId(ownerId string) ([]models.Company, error)
	// company role
	CreateCompanyRoleModel(id string, entity models.CreateCompanyRoleModel) error
	GetRolesListByCompany(companyID string) ([]models.CompanyRoleListByCompany, error)
	//orders
	CreateOrderModel(entity models.CreateOrderModel) error
	GetOrdersList(companyID string) ([]models.OrderList, error)
}
