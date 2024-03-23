package storage

import (
	"bw-erp/storage/postgres"
	"bw-erp/storage/repo"

	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	// user
	// CreateUserModel(id string, entity models.CreateUserModel) error
	// GetUserByPhone(phone string) (models.AuthUserModel, error)
	// GetUserById(id string) (models.User, error)
	// GetUsersList(companyID string) ([]models.User, error)
	// ChangeUserPassword(userID string, entity models.ChangePasswordRequest) error

	// company
	// CreateCompanyModel(id string, entity models.CreateCompanyModel) error
	// GetCompanyByOwnerId(ownerId string) ([]models.Company, error)

	// role
	// CreateRoleModel(id string, entity models.CreateRoleModel) error
	// GetRolesListByCompany(companyID string) ([]models.RoleListByCompany, error)
	// GetPermissionsToRole(models.GetPermissionToRoleRequest) error
	// GetRoleByPrimaryKey(roleID string) (models.RoleByPrimaryKey, error)

	//orders
	// CreateOrderModel(entity models.CreateOrderModel) (id int, err error)
	// GetOrdersList(companyID string, queryParam models.OrdersListRequest) (res models.OrderListResponse, err error)
	// GetOrderLocation(ID int) (models.Order, error)
	// GetOrderDetailedByPrimaryKey(ID int) (models.OrderShowResponse, error)
	// UpdateOrder(entity *models.UpdateOrderRequest) (rowsAffected int64, err error)
	// GetOrdersByStatus(companyID string, Status int) (order []models.Order, err error)
	// GetOrderByPhone(companyID string, Phone string) (models.Order, error)
	// GetOrderByPrimaryKey(ID int) (models.OrderShowResponse, error)

	//order-items
	// CreateOrderItemModel(entity models.CreateOrderItemModel) error
	// UpdateOrderItemModel(entity models.UpdateOrderItemRequest) (rowsAffected int64, err error)

	//order item type
	// CreateOrderItemTypeModel(id string, entity models.OrderItemTypeModel) error
	// GetOrderItemTypesByCompany(CompanyID string) ([]models.OrderItemByCompany, error)
	// UpdateOrderItemTypeModel(entity models.EditOrderItemTypeRequest) (rowsAffected int64, err error)

	//company bots
	// CreateCompanyBotModel(CompanyID string, entity models.CreateCompanyBotModel) error
	// GetTelegramBotByCompany(CompanyID string) (models.CompanyTelegramBot, error)
	// GetTelegramOrderBot() ([]models.CompanyTelegramBot, error)

	// bot-users
	// GetBotUserByChatIDModel(ChatID int64, BotID int64) (models.BotUser, error)
	// CreateBotUserModel(entity models.CreateBotUserModel) error
	// GetSelectedUser(BotID int64, Phone string) (models.SelectedUser, error)
	// UpdateBotUserModel(entity models.BotUser) (rowsAffected int64, err error)
	// GetBotUserByCompany(BotID int64, ChatID int64) (botUser models.BotUserByCompany, err error)
	// GetBotUserByUserID(UserID string) (models.BotUser, error)
	// GetNotificationGroup(CompanyID string) (models.BotUserByCompany, error)

	// telegram-session
	// GetTelegramSessionByChatIDBotID(ChatID int64, BotID int64) (models.TelegramSessionModel, error)
	// DeleteTelegramSession(ID int) (rowsAffected int64, err error)
	// CreateTelegramSessionModel(entity models.TelegramSessionModel) error

	//work volume
	// GetWorkVolumeList(companyID string) ([]models.WorkVolume, error)

	// permission
	// GetPermissionList(Scope string) ([]models.Permission, error)
	// GetPermissionByPrimaryKey(ID string) (models.Permission, error)

	//clients
	// CreateClientModel(entity models.CreateClientModel) (id int, err error)
	// GetClientsList(companyID string, queryParam models.ClientListRequest) (res models.ClientListResponse, err error)
	// GetClientByPrimaryKey(ID int) (models.GetClientByPrimaryKeyResponse, error)
	// UpdateClient(entity *models.UpdateClientRequest) (rowsAffected int64, err error)

	//telegram groups
	// CreateTelegramGroupModel(entity models.CreateTelegramGroupRequest) error
	// GetNotificationGroup(CompanyID string, Status int) (models.TelegramGroup, error)

	Client() repo.ClientStorageI
	BotUser() repo.BotUserI
	Company() repo.CompanyStorageI
	OrderItemType() repo.OrderItemTypeI
	OrderItem() repo.OrderItemI
	Order() repo.OrderI
	Permission() repo.PermissionI
	Role() repo.RoleI
	Statistics() repo.StatisticsI
	TelegramBot() repo.TelegramBotI
	TelegramGroup() repo.TelegramGroupI
	TelegramSession() repo.TelegramSessionI
	User() repo.UserI
}

type storagePg struct {
	db          *sqlx.DB
	clientRepo  repo.ClientStorageI
	botUserRepo repo.BotUserI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		db:         db,
		clientRepo: postgres.NewClientRepo(db),
	}
}

func (s storagePg) Client() repo.ClientStorageI {
	return s.clientRepo
}

func (s storagePg) BotUser() repo.BotUserI {
	return s.botUserRepo
}

func (s storagePg) Company() repo.CompanyStorageI {
	return s.Company()
}

func (s storagePg) OrderItemType() repo.OrderItemTypeI {
	return s.OrderItemType()
}

func (s storagePg) OrderItem() repo.OrderItemI {
	return s.OrderItem()
}

func (s storagePg) Order() repo.OrderI {
	return s.Order()
}

func (s storagePg) Permission() repo.PermissionI {
	return s.Permission()
}

func (s storagePg) Role() repo.RoleI {
	return s.Role()
}

func (s storagePg) Statistics() repo.StatisticsI {
	return s.Statistics()
}

func(s storagePg) TelegramBot() repo.TelegramBotI {
	return s.TelegramBot()
}

func(s storagePg) TelegramGroup() repo.TelegramGroupI {
	return s.TelegramGroup()
}

func(s storagePg) TelegramSession() repo.TelegramSessionI {
	return s.TelegramSession()
}

func(s storagePg) User() repo.UserI {
	return s.User()
}