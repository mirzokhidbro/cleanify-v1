package handlers

import (
	"bw-erp/api/http"
	"bw-erp/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateOrderItemModel(c *gin.Context) {
	var body models.CreateOrderItemModel
	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	_, err := h.Stg.Order().GetByPrimaryKey(body.OrderID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, "Order not found")
		return
	}

	orderItemType, err := h.Stg.OrderItemType().GetById(body.OrderItemTypeID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	body.Price = orderItemType.Price
	body.ItemType = orderItemType.Name

	err = h.Stg.OrderItem().Create(body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.Created, "Created successfully!")
}

func (h *Handler) UpdateOrderItemModel(c *gin.Context) {
	var body models.UpdateOrderItemRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	_, err := h.Stg.OrderItem().Update(body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, "Update successfully!")
}
