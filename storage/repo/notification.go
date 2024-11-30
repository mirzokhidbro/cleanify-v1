package repo

import "bw-erp/models"

type NotificationI interface {
	Create(entity models.CreateNotificationModel) (notification_id int, err error)
	GetMyNotifications(entity models.GetMyNotificationsRequest) ([]models.GetMyNotificationsResponse, error)
	GetMyLatestNotifications(entity models.GetMyNotificationsRequest) (models.GetMyNotificationsResponse, error)
	GetNotificationsByID(entity models.GetNotificationsByIDRequest) ([]models.GetMyNotificationsResponse, error)
	GetUnreadNotificationsCount(userID string) (int, error)
}
