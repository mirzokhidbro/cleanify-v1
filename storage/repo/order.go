package repo

import "bw-erp/models"

type OrderI interface {
	Create(entity models.CreateOrderModel) (id int, err error)
	GetList(companyID string, queryParam models.OrdersListRequest) (res models.OrderListResponse, err error)
	GetLocation(ID int) (models.Order, error)
	GetDetailedByPrimaryKey(ID int) (models.OrderShowResponse, error)
	Update(entity *models.UpdateOrderRequest) (rowsAffected int64, err error)
	GetByStatus(companyID string, Status int) (order []models.Order, err error)
	GetByPhone(companyID string, Phone string) (models.Order, error)
	GetByPrimaryKey(ID int) (models.OrderShowResponse, error)
	Delete(entity models.DeleteOrderRequest) error
}
