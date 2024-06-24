package repo

import "bw-erp/models"

type ClientStorageI interface {
	Create(entity models.CreateClientModel) (id int, err error)
	GetList(companyID string, queryParam models.ClientListRequest) (res models.ClientListResponse, err error)
	GetByPrimaryKey(ID int) (models.GetClientByPrimaryKeyResponse, error)
	GetByPhoneNumber(Phone string) (models.Client, error)
	Update(companyID string, entity models.UpdateClientRequest) (rowsAffected int64, err error)
}
