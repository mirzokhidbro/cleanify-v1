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

	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	id := uuid.New()

	err := h.Stg.User().Create(id.String(), body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, id)
}

func (h *Handler) GetList(c *gin.Context) {
	var body models.GetUserListRequest
	if err := c.ShouldBindQuery(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	users, err := h.Stg.User().GetList(body.CompanyID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	h.handleResponse(c, http.OK, users)
}

func (h *Handler) Edit(c *gin.Context) {
	var body models.UpdateUserRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	if !utils.IsValidUUID(body.ID) {
		h.handleResponse(c, http.InvalidArgument, "user id is the invalid uuid")
		return
	}

	_, err := h.Stg.User().Edit(body)

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

func (h *Handler) GetCouriesList(c *gin.Context) {
	var req models.GetCouriesListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	couriers, err := h.Stg.User().GetCouriesList(req.CompanyID)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, couriers)
}
