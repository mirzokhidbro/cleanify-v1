package repo

import "bw-erp/models"

type StatusChangeHistoryI interface {
	Create(entity models.CreateStatusChangeHistoryModel) error
}
