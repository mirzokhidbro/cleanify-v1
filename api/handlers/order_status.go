package handlers

import (
	"bw-erp/api/http"
	"bw-erp/models"
	"bw-erp/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetOrderStatusesList(c *gin.Context) {
	var body models.GetOrderStatusListRequest
	if err := c.ShouldBindQuery(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	if !utils.IsValidUUID(body.CompanyID) {
		h.handleResponse(c, http.InvalidArgument, "company id is an invalid uuid")
		return
	}

	statuses, err := h.Stg.OrderStatus().GetList(body.CompanyID)

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

func (h *Handler) GetOrderStatusById(c *gin.Context) {
	ID := c.Param("id")
	statusID, err := strconv.Atoi(ID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	data, err := h.Stg.OrderStatus().GetById(statusID)

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, data)
}
