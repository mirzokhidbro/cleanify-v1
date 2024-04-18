package handlers

import (
	"bw-erp/api/http"
	"bw-erp/models"
	"bw-erp/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) Create(c *gin.Context) {
	var body models.CreateUserModel

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

	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	id := uuid.New()

	err = h.Stg.User().Create(id.String(), body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, id)
}

func (h *Handler) GetList(c *gin.Context) {
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
	users, err := h.Stg.User().GetList(*user.CompanyID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	h.handleResponse(c, http.OK, users)
}

func (h *Handler) Edit(c *gin.Context) {
	var body models.UpdateUserRequest
	companyID := c.Param("company-id")
	if !utils.IsValidUUID(companyID) {
		h.handleResponse(c, http.InvalidArgument, "company id is the invalid uuid")
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	if !utils.IsValidUUID(body.ID) {
		h.handleResponse(c, http.InvalidArgument, "user id is the invalid uuid")
		return
	}

	_, err := h.Stg.User().Edit(companyID, body)

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, true)
}

func (h *Handler) GetById(c *gin.Context) {
	userID := c.Param("user-id")

	if !utils.IsValidUUID(userID) {
		h.handleResponse(c, http.InvalidArgument, "user id is an invalid uuid")
		return
	}

	user, err := h.Stg.User().GetById(userID)

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, user)
}
