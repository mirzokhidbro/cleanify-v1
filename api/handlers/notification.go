package handlers

import (
	"bw-erp/api/http"
	"bw-erp/models"

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
