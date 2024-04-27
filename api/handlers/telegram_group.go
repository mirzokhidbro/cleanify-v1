package handlers

import (
	"bw-erp/api/http"
	"bw-erp/models"
	"bw-erp/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) VerificationGroup(c *gin.Context) {
	var body models.GroupVerificationRequest
	companyID := c.Param("company-id")

	if !utils.IsValidUUID(companyID) {
		h.handleResponse(c, http.InvalidArgument, "company id is an invalid uuid")
		return
	}

	_, err := h.Stg.Company().GetById(companyID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, "company not found")
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	_, err = h.Stg.TelegramGroup().Verification(body.Code, companyID)

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.Created, "Success!")
}

func (h *Handler) GetTelegramGroupList(c *gin.Context) {
	companyID := c.Param("company-id")

	if !utils.IsValidUUID(companyID) {
		h.handleResponse(c, http.InvalidArgument, "company id is an invalid uuid")
		return
	}

	_, err := h.Stg.Company().GetById(companyID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, "company not found")
		return
	}

	data, err := h.Stg.TelegramGroup().GetList(companyID)

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.Created, data)
}

func (h *Handler) GetTelegramGroupByPrimaryKey(c *gin.Context) {
	ID := c.Param("id")
	groupID, err := strconv.Atoi(ID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	data, err := h.Stg.TelegramGroup().GetByPrimaryKey(groupID)

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, data)
}

func (h *Handler) UpdateTelegramGroup(c *gin.Context) {
	ID := c.Param("id")
	var body models.TelegramGroupEditRequest
	groupID, err := strconv.Atoi(ID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	data, err := h.Stg.TelegramGroup().Update(groupID, body)

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, data)
}
