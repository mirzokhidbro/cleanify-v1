package handlers

import (
	"bw-erp/api/http"
	"bw-erp/models"
	"bw-erp/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateRoleModel(c *gin.Context) {
	var body models.CreateRoleModel
	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	_, err := h.Stg.Company().GetById(body.CompanyId)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	// token, err := utils.ExtractTokenID(c)

	// if err != nil {
	// 	h.handleResponse(c, http.BadRequest, err.Error())
	// 	return
	// }

	// user, err := h.Stg.GetUserById(token.UserID)
	// if err != nil {
	// 	h.handleResponse(c, http.BadRequest, err.Error())
	// 	return
	// }

	// body.CompanyId = *user.CompanyID
	id := uuid.New()

	err = h.Stg.Role().Create(id.String(), body)
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

	token, err := utils.ExtractTokenID(c)

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	_, err = h.Stg.User().GetById(token.UserID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	data, err := h.Stg.Role().GetListByCompany(companyID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, data)
}

func (h *Handler) GetPermissionsToRole(c *gin.Context) {
	var body models.GetPermissionToRoleRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	token, err := utils.ExtractTokenID(c)

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	user, err := h.Stg.User().GetById(token.UserID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	body.CompanyID = *user.CompanyID

	err = h.Stg.Role().GetPermissionsToRole(body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, "Success!")
}

func (h *Handler) GetRoleByPrimaryKey(c *gin.Context) {
	roleID := c.Param("role-id")
	if !utils.IsValidUUID(roleID) {
		h.handleResponse(c, http.InvalidArgument, "role id is an invalid uuid")
		return
	}

	data, err := h.Stg.Role().GetByPrimaryKey(roleID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, data)
}
