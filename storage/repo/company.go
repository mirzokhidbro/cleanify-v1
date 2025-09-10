package repo

import "bw-erp/models"

type CompanyStorageI interface {
	Create(id string, entity models.CreateCompanyModel) error
	GetByOwnerId(ownerId int64) ([]models.Company, error)
	GetById(id string) (models.Company, error)
}
