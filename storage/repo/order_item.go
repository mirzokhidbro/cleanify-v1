package repo

import "bw-erp/models"

type OrderItemI interface {
	Create(userID int64, entity models.CreateOrderItemModel) error
	Update(entity models.UpdateOrderItemRequest) (rowsAffected int64, err error)
	DeleteByID(ID int) error
	UpdateStatus(userID int64, entity models.UpdateOrderItemStatusRequest) (rowsAffected int64, err error)
}
