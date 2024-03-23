package repo

import "bw-erp/models"

type OrderItemTypeI interface {
	Create(id string, entity models.OrderItemTypeModel) error
	GetByCompany(CompanyID string) ([]models.OrderItemByCompany, error)
	Update(entity models.EditOrderItemTypeRequest) (rowsAffected int64, err error)
}
