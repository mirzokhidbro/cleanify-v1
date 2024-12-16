package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	webpush "github.com/SherClockHolmes/webpush-go"
)

var (
	vapidPrivateKey = os.Getenv("2XVD9yMwyRaFdvliK_2zxTu2pvW1Wv08h3HXm15g0hg")
	vapidPublicKey  = os.Getenv("BHcLciJSf0BX5QG9nHyRTBkCGX9FikeyoagyPGTq1DoEf8TR1qYgJQPaosv8SruBpBYak9QNiUIpXaJshrSwUv4")
)

type PushSubscription struct {
	Endpoint       string               `json:"endpoint"`
	ExpirationTime int64                `json:"expirationTime"`
	Keys           PushSubscriptionKeys `json:"keys"`
}

type PushSubscriptionKeys struct {
	P256dh string `json:"p256dh"`
	Auth   string `json:"auth"`
}

type WebPushManager struct {
	subscriptions map[string][]PushSubscription // userID -> subscriptions
}

var webPushManager *WebPushManager

func InitWebPushManager() {
	webPushManager = &WebPushManager{
		subscriptions: make(map[string][]PushSubscription),
	}
}

func GetWebPushManager() *WebPushManager {
	if webPushManager == nil {
		InitWebPushManager()
	}
	return webPushManager
}

func (m *WebPushManager) SaveSubscription(userID string, subscription PushSubscription) {
	if _, exists := m.subscriptions[userID]; !exists {
		m.subscriptions[userID] = []PushSubscription{}
	}
	m.subscriptions[userID] = append(m.subscriptions[userID], subscription)
}

func (m *WebPushManager) SendNotification(userID string, payload interface{}) error {
	subs, exists := m.subscriptions[userID]
	if !exists {
		return fmt.Errorf("no subscriptions found for user %s", userID)
	}

	message, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal notification: %v", err)
	}

	for _, sub := range subs {
		s := &webpush.Subscription{
			Endpoint: sub.Endpoint,
			Keys: webpush.Keys{
				P256dh: sub.Keys.P256dh,
				Auth:   sub.Keys.Auth,
			},
		}

		resp, err := webpush.SendNotification(message, s, &webpush.Options{
			Subscriber:      "example@example.com", // Change this
			VAPIDPublicKey:  vapidPublicKey,
			VAPIDPrivateKey: vapidPrivateKey,
			TTL:             30,
		})
		if err != nil {
			log.Printf("Failed to send notification to subscription: %v", err)
			continue
		}
		resp.Body.Close()
	}
	return nil
}
