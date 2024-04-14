package handlers

import (
	"bw-erp/api/http"
	"bw-erp/models"
	"bw-erp/pkg/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetOrderStatusesList(c *gin.Context) {
	companyID := c.Param("company-id")
	if !utils.IsValidUUID(companyID) {
		h.handleResponse(c, http.InvalidArgument, "company id is an invalid uuid")
		return
	}

	statuses, err := h.Stg.OrderStatus().GetList(companyID)

	if err != nil {
		h.handleResponse(c, http.OK, err.Error())
		return
	}

	h.handleResponse(c, http.OK, statuses)
}

func (h *Handler) UpdateOrderStatusModel(c *gin.Context) {
	var body models.UpdateOrderStatusRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	_, err := h.Stg.OrderStatus().Update(body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, "Update successfully!")
}
