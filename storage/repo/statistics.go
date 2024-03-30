package repo

import "bw-erp/models"

type StatisticsI interface {
	GetWorkVolume(companyID string) ([]models.WorkVolume, error)
}
