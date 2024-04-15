package repo

import "bw-erp/models"

type OrderItemI interface {
	Create(entity models.CreateOrderItemModel) error
	Update(entity models.UpdateOrderItemRequest) (rowsAffected int64, err error)
	DeleteByID(ID int) error
}
