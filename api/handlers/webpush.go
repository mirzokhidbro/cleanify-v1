package handlers

import (
	"bw-erp/api/http"
	"bw-erp/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SavePushSubscription(c *gin.Context) {
	fmt.Println("validate")
	var subscription models.CreatePushSubscriptionRequest
	if err := c.ShouldBindJSON(&subscription); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	_, err := h.Stg.WebPush().CreatePushSubscription(subscription)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, "OK")
}

// GetPushSubscription returns subscription details for a user
// func (h *Handler) GetPushSubscription(c *gin.Context) {
// 	userID := c.Query("user_id")
// 	if userID == "" {
// 		h.handleResponse(c, http.BadRequest, "user_id is required")
// 		return
// 	}

// 	subscription, err := h.Stg.WebPush().GetPushSubscription(userID)
// 	if err != nil {
// 		h.handleResponse(c, http.InternalServerError, err.Error())
// 		return
// 	}

// 	if subscription == nil {
// 		h.handleResponse(c, http.NOT_FOUND, "subscription not found")
// 		return
// 	}

// 	h.handleResponse(c, http.OK, subscription)
// }

// // DeletePushSubscription removes a user's push subscription
// func (h *Handler) DeletePushSubscription(c *gin.Context) {
// 	userID := c.Query("user_id")
// 	if userID == "" {
// 		h.handleResponse(c, http.BadRequest, "user_id is required")
// 		return
// 	}

// 	err := h.Stg.WebPush().DeletePushSubscription(userID)
// 	if err != nil {
// 		h.handleResponse(c, http.InternalServerError, err.Error())
// 		return
// 	}

// 	h.handleResponse(c, http.OK, "subscription deleted successfully")
// }

// // GetAllPushSubscriptions returns all push subscriptions
// func (h *Handler) GetAllPushSubscriptions(c *gin.Context) {
// 	var params models.GetPushSubscriptionResponse
// 	subscriptions, err := h.Stg.WebPush().GetAllPushSubscriptions(&params)
// 	if err != nil {
// 		h.handleResponse(c, http.InternalServerError, err.Error())
// 		return
// 	}

// 	h.handleResponse(c, http.OK, subscriptions)
// }

// // SendPushNotification sends a push notification to a specific user
// func (h *Handler) SendPushNotification(c *gin.Context) {
// 	userID := c.Query("user_id")
// 	if userID == "" {
// 		h.handleResponse(c, http.BadRequest, "user_id is required")
// 		return
// 	}

// 	var notification map[string]interface{}
// 	if err := c.ShouldBindJSON(&notification); err != nil {
// 		h.handleResponse(c, http.BadRequest, err.Error())
// 		return
// 	}

// 	subscription, err := h.Stg.WebPush().GetPushSubscription(userID)
// 	if err != nil {
// 		h.handleResponse(c, http.InternalServerError, err.Error())
// 		return
// 	}

// 	if subscription == nil {
// 		h.handleResponse(c, http.NOT_FOUND, "subscription not found")
// 		return
// 	}

// 	// TODO: Implement actual push notification sending using web-push library
// 	// This will be implemented in the next step with VAPID keys

// 	h.handleResponse(c, http.OK, "notification sent successfully")
// }
