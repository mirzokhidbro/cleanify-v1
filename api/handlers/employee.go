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

	data, err := h.Stg.Employee().GetList(body)
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

func (h *Handler) AddTransaction(c *gin.Context) {
	var body models.EmployeeTransactionRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	if body.Salary == 0 && body.ReceivedMoney == 0 {
		h.handleResponse(c, http.InvalidArgument, "At least one of received money and salary should be sent")
		return
	}

	err := utils.TokenValid(c)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	jwtData, err := utils.ExtractTokenID(c)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	body.UserID = jwtData.UserID

	employee, _ := h.Stg.Employee().GetDetailedData(models.ShowEmployeeRequest{
		CompanyID:  body.CompanyID,
		EmployeeID: body.EmployeeID,
	})

	if employee.ID == 0 {
		h.handleResponse(c, http.BadRequest, "Employee not found")
		return
	}

	err = h.Stg.Employee().AddTransaction(body)

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, "OK!")

}

func (h *Handler) Attendance(c *gin.Context) {
	var body models.AttendanceEmployeeRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	if !utils.IsValidUUID(body.CompanyID) {
		h.handleResponse(c, http.InvalidArgument, "company id is an invalid uuid")
		return
	}

	err := h.Stg.Employee().Attendance(body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.Created, "OK!")
}
