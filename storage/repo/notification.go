package repo

import "bw-erp/models"

type NotificationI interface {
	GetMyNotifications(entity models.GetMyNotificationsRequest) ([]models.GetMyNotificationsResponse, error)
}
