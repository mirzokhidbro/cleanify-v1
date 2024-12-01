package handlers

import (
	"bw-erp/api/http"
	"bw-erp/models"
	"bw-erp/pkg/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetMyNotifications(c *gin.Context) {
	var body models.GetMyNotificationsRequest

	if err := c.ShouldBindQuery(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	data, err := h.Stg.Notification().GetMyNotifications(body)

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, data)
}

func (h *Handler) HandleNotificationWebSocket(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		h.handleResponse(c, http.BadRequest, "user_id is required")
		return
	}

	if err := utils.GetManager().HandleConnection(c, userID); err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
}

func (h *Handler) SendNotificationToUser(userID string, notification models.GetMyNotificationsResponse) error {
	return utils.GetManager().SendMessage(userID, notification)
}

func (h *Handler) UnreadNotificationsCount(c *gin.Context) {
	token, err := utils.ExtractTokenID(c)

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	user, err := h.Stg.User().GetById(token.UserID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	count, err := h.Stg.Notification().GetUnreadNotificationsCount(user.ID)

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, count)
}
