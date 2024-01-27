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
	GetOrderByPrimaryKey(ID int) (models.Order, error)
	UpdateOrder(entity *models.UpdateOrderRequest) (rowsAffected int64, err error)

	//order-items
	CreateOrderItemModel(entity models.CreateOrderItemModel) error

	//order item type
	CreateOrderItemTypeModel(id string, entity models.OrderItemTypeModel) error
	GetOrderItemTypesByCompany(CompanyID string) ([]models.OrderItemByCompany, error)

	//company bots
	CreateCompanyBotModel(CompanyID string, entity models.CreateCompanyBotModel) error
	GetTelegramBotByCompany(CompanyID string) (models.CompanyTelegramBot, error)
	GetTelegramOrderBot() ([]models.CompanyTelegramBot, error)

	// bot-users
	GetBotUserByChatIDModel(ChatID int64) (models.BotUser, error)
	CreateBotUserModel(entity models.CreateBotUserModel) error
	GetSelectedUser(BotID int64, Phone string) (models.SelectedUser, error)
	UpdateBotUserModel(entity models.BotUser) (rowsAffected int64, err error)
}
