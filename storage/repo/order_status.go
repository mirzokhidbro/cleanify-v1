package repo

import "bw-erp/models"

type OrderStatusI interface {
	GetList(CompanyID string) ([]models.OrderStatusListResponse, error)
	Update(entity models.UpdateOrderStatusRequest) (rowsAffected int64, err error)
}
