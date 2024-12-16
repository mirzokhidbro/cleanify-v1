package repo

import "bw-erp/models"

type WebPushStorageI interface {
	CreatePushSubscription(models.CreatePushSubscriptionRequest) (int64, error)
	GetPushSubscription(userID string) (*models.PushSubscription, error)
	DeletePushSubscription(userID string) error
	GetAllPushSubscriptions(*models.GetPushSubscriptionResponse) (*models.GetPushSubscriptionResponse, error)
}
