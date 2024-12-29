package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	webpush "github.com/SherClockHolmes/webpush-go"
)

var (
	vapidPublicKey  = "BHcLciJSf0BX5QG9nHyRTBkCGX9FikeyoagyPGTq1DoEf8TR1qYgJQPaosv8SruBpBYak9QNiUIpXaJshrSwUv4"
	vapidPrivateKey = "2XVD9yMwyRaFdvliK_2zxTu2pvW1Wv08h3HXm15g0hg"
	vapidSubject    = "mailto:mirzoxidadxamjanov01@gmail.com"
)

type PushSubscriptionKeys struct {
	Auth   string `json:"auth"`
	P256dh string `json:"p256dh"`
}

type PushSubscriptionData struct {
	Endpoint       string               `json:"endpoint"`
	ExpirationTime *int64               `json:"expirationTime"`
	Keys           PushSubscriptionKeys `json:"keys"`
}

type BrowserSubscription struct {
	UserAgent    string               `json:"userAgent"`
	Subscription PushSubscriptionData `json:"subscription"`
}

func SendPushNotification(subscription []byte, message string) error {
	var browserSub BrowserSubscription
	if err := json.Unmarshal(subscription, &browserSub); err != nil {
		return fmt.Errorf("failed to unmarshal subscription: %v", err)
	}

	s := &webpush.Subscription{
		Endpoint: browserSub.Subscription.Endpoint,
		Keys: webpush.Keys{
			Auth:   browserSub.Subscription.Keys.Auth,
			P256dh: browserSub.Subscription.Keys.P256dh,
		},
	}

	resp, err := webpush.SendNotification([]byte(message), s, &webpush.Options{
		Subscriber:      vapidSubject,
		VAPIDPublicKey:  vapidPublicKey,
		VAPIDPrivateKey: vapidPrivateKey,
		TTL:             30,
	})
	if err != nil {
		return fmt.Errorf("failed to send notification: %v", err)
	}
	defer resp.Body.Close()

	// Response ni o'qish va log ga chiqarish
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
	} else {
		log.Printf("Response body: %s", string(body))
	}

	log.Printf("Push notification response: Status=%s, StatusCode=%d, Headers=%v",
		resp.Status, resp.StatusCode, resp.Header)

	return nil
}
