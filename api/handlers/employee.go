package handlers

import (
	"bw-erp/api/http"
	"bw-erp/models"
	"bw-erp/pkg/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateEmployee(c *gin.Context) {
	var body models.CreateEmployeeRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	if !utils.IsValidUUID(body.CompanyID) {
		h.handleResponse(c, http.InvalidArgument, "company id is an invalid uuid")
		return
	}

	err := h.Stg.Employee().Create(body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.Created, "OK!")
}

func (h *Handler) GetEmployeeList(c *gin.Context) {
	var body models.GetEmployeeListRequest
	if err := c.ShouldBindQuery(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	data, err := h.Stg.Employee().GetList(body.CompanyID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, data)
}

func (h *Handler) ShowEmployeeDetailedData(c *gin.Context) {
	var body models.ShowEmployeeRequest
	if err := c.ShouldBindQuery(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	data, err := h.Stg.Employee().GetDetailedData(body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, data)
}
