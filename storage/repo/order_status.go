package repo

import "bw-erp/models"

type OrderStatusI interface {
	GetList(CompanyID string) ([]models.OrderStatus, error)
	Update(entity models.UpdateOrderStatusRequest) (rowsAffected int64, err error)
	GetById(id int) (models.OrderStatus, error)
}
