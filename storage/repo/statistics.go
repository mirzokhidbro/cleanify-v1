package repo

import "bw-erp/models"

type StatisticsI interface {
	GetWorkVolume(companyID string) ([]models.WorkVolume, error)
	GetServicePaymentStatistics(entity models.GetServicePaymentStatisticsRequest) ([]models.ServicePaymentStatistics, error)
}
