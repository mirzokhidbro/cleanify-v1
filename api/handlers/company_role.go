package handlers

import (
	"bw-erp/api/http"
	"bw-erp/models"
	"bw-erp/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateCompanyRoleModel(c *gin.Context) {
	var body models.CreateCompanyRoleModel
	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	id := uuid.New()

	err := h.Stg.CreateCompanyRoleModel(id.String(), body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.Created, id)
}

func (h *Handler) GetRolesListByCompany(c *gin.Context) {
	companyID := c.Param("company-id")
	if !utils.IsValidUUID(companyID) {
		h.handleResponse(c, http.InvalidArgument, "company id is an invalid uuid")
		return
	}

	data, err := h.Stg.GetRolesListByCompany(companyID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, data)
}
