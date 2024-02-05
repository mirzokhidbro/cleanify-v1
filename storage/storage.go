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
	ChangeUserPassword(userID string, entity models.ChangePasswordRequest) error

	// company
	CreateCompanyModel(id string, entity models.CreateCompanyModel) error
	GetCompanyByOwnerId(ownerId string) ([]models.Company, error)

	// company role
	CreateRoleModel(id string, entity models.CreateRoleModel) error
	GetRolesListByCompany(companyID string) ([]models.RoleListByCompany, error)
	GetPermissionsToRole(models.GetPermissionToRoleRequest) error

	//orders
	CreateOrderModel(entity models.CreateOrderModel) (id int, err error)
	GetOrdersList(companyID string, queryParam models.OrdersListRequest) (res models.OrderListResponse, err error)
	GetOrderLocation(ID int) (models.Order, error)
	GetOrderByPrimaryKey(ID int) (models.Order, error)
	UpdateOrder(entity *models.UpdateOrderRequest) (rowsAffected int64, err error)

	//order-items
	CreateOrderItemModel(entity models.CreateOrderItemModel) error
	UpdateOrderItemModel(entity models.UpdateOrderItemRequest) (rowsAffected int64, err error)

	//order item type
	CreateOrderItemTypeModel(id string, entity models.OrderItemTypeModel) error
	GetOrderItemTypesByCompany(CompanyID string) ([]models.OrderItemByCompany, error)
	UpdateOrderItemTypeModel(entity models.EditOrderItemTypeRequest) (rowsAffected int64, err error)

	//company bots
	CreateCompanyBotModel(CompanyID string, entity models.CreateCompanyBotModel) error
	GetTelegramBotByCompany(CompanyID string) (models.CompanyTelegramBot, error)
	GetTelegramOrderBot() ([]models.CompanyTelegramBot, error)

	// bot-users
	GetBotUserByChatIDModel(ChatID int64, BotID int64) (models.BotUser, error)
	CreateBotUserModel(entity models.CreateBotUserModel) error
	GetSelectedUser(BotID int64, Phone string) (models.SelectedUser, error)
	UpdateBotUserModel(entity models.BotUser) (rowsAffected int64, err error)
	GetBotUserByCompany(BotID int64, ChatID int64) (botUser models.BotUserByCompany, err error)
	GetBotUserByUserID(UserID string) (models.BotUser, error)

	// telegram-session
	GetTelegramSessionByChatIDBotID(ChatID int64, BotID int64) (models.TelegramSessionModel, error)
	DeleteTelegramSession(ID int) (rowsAffected int64, err error)
	CreateTelegramSessionModel(entity models.TelegramSessionModel) error

	//work volume
	GetWorkVolumeList(companyID string) ([]models.WorkVolume, error)

	// permission
	GetPermissionList(Scope string) ([]models.Permission, error)
	GetPermissionByPrimaryKey(ID string) (models.Permission, error)
}
