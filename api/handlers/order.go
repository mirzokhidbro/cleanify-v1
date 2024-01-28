package handlers

import (
	"bw-erp/api/http"
	"bw-erp/models"
	"bw-erp/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateOrderModel(c *gin.Context) {
	var body models.CreateOrderModel
	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	_, err := h.Stg.CreateOrderModel(body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.Created, "Created successfully!")
}

func (h *Handler) GetOrdersList(c *gin.Context) {
	companyID := c.Param("company-id")
	if !utils.IsValidUUID(companyID) {
		h.handleResponse(c, http.InvalidArgument, "company id is an invalid uuid")
		return
	}
	data, err := h.Stg.GetOrdersList(companyID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, data)
}

func (h *Handler) GetOrderByPrimaryKey(c *gin.Context) {
	orderID := c.Param("order-id")
	orderId, err := strconv.Atoi(orderID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	data, err := h.Stg.GetOrderByPrimaryKey(orderId)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	h.handleResponse(c, http.OK, data)
}

func (h *Handler) UpdateOrderModel(c *gin.Context) {
	var body models.UpdateOrderRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	rowsAffected, err := h.Stg.UpdateOrder(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, rowsAffected)

}
