package handlers

import (
	"bw-erp/api/http"
	"bw-erp/pkg/utils"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Ping(c *gin.Context) {
	// Yangi subscription data
	testSubscription := `{"userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36", "subscription": {"keys": {"auth": "O3O2JoSrTYq2nnngMJjYsw", "p256dh": "BLmoq08Khm0KHY0w-rJh2nERX1OchTu3autVMubNQLRk9wCGiprzVd_LyE_LyQX40icJYTrfhgF78ie208c0vfo"}, "endpoint": "https://fcm.googleapis.com/fcm/send/exJXVs3D20Y:APA91bGxmdhuhlfXztsr4dzRVOQ5163naVOgmG6svR11FxMK9vaPk9cJL838AxkqIdqOyojlpJ1l8ipaeFrUNMjHRu0ON5c7awykpCKFKY2YuMIYwd4S4lUmzZSqqQGR6TA2zWYWGjQh", "expirationTime": null}}`

	// Push notification payload
	payload := map[string]interface{}{
		"title": "BW-ERP Notification",
		"body":  "Yangi xabar keldi!",
		"icon":  "/icon.png",
		"data": map[string]interface{}{
			"url": "/dashboard",
		},
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	// Send test notification
	err = utils.SendPushNotification([]byte(testSubscription), string(payloadJSON))
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, "pong")
}
