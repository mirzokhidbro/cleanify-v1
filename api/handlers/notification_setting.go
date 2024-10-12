package handlers

import (
	"bw-erp/api/http"
	"bw-erp/models"
	"bw-erp/pkg/utils"

	"github.com/gin-gonic/gin"
)

func (h Handler) SetNotificationSetting(c *gin.Context) {
	var body models.SetNotificationSettingRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	if !utils.IsValidUUID(body.CompanyID) {
		h.handleResponse(c, http.InvalidArgument, "company id is an invalid uuid")
		return
	}

	err := h.Stg.NotificationSetting().NotificationSetting(body)

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, "Created successfully")
}

func (h Handler) UsersListForNotificationSettings(c *gin.Context) {
	var body models.UsersListForNotificationSettingsRequest

	if err := c.ShouldBindQuery(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	if !utils.IsValidUUID(body.CompanyID) {
		h.handleResponse(c, http.InvalidArgument, "company id is an invalid uuid")
		return
	}

	data := h.Stg.NotificationSetting().UsersListForNotificationSettings(body.CompanyID)
	h.handleResponse(c, http.OK, data)
}

func (h *Handler) GetUsersByStatus(c *gin.Context) {
	var body models.GetUsersByStatusRequest

	if err := c.ShouldBindQuery(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	if !utils.IsValidUUID(body.CompanyID) {
		h.handleResponse(c, http.InvalidArgument, "company id is an invalid uuid")
		return
	}

	data, err := h.Stg.NotificationSetting().GetUsersByStatus(body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, data)
}
