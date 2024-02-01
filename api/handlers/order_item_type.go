package handlers

import (
	"bw-erp/api/http"
	"bw-erp/models"
	"bw-erp/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateOrderItemTypeModel(c *gin.Context) {
	var body models.OrderItemTypeModel

	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	if !utils.IsValidUUID(body.CopmanyID) {
		h.handleResponse(c, http.InvalidArgument, "company id is an invalid uuid")
		return
	}

	id := uuid.New()

	err := h.Stg.CreateOrderItemTypeModel(id.String(), body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.Created, id)
}

func (h *Handler) GetOrderItemTypesByCompany(c *gin.Context) {
	// companyID := c.Param("company-id")
	// if !utils.IsValidUUID(companyID) {
	// 	h.handleResponse(c, http.InvalidArgument, "company id is an invalid uuid")
	// 	return
	// }

	token, err := utils.ExtractTokenID(c)

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	user, err := h.Stg.GetUserById(token.UserID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	data, err := h.Stg.GetOrderItemTypesByCompany(*user.CompanyID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, data)
}
