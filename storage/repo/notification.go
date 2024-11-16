package repo

import "bw-erp/models"

type NotificationI interface {
	GetMyNotifications(entity models.GetMyNotificationsRequest) ([]models.GetMyNotificationsResponse, error)
	GetMyLatestNotifications(entity models.GetMyNotificationsRequest) (models.GetMyNotificationsResponse, error)
	GetNotificationsByStatus(entity models.GetNotificationsByStatusRequest) ([]models.GetMyNotificationsResponse, error)
	GetUnreadNotificationsCount(userID string, companyID string) (int, error)
}
