package handlers

import (
	"bw-erp/api/http"
	"bw-erp/models"
	"bw-erp/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateCompanyModel(c *gin.Context) {
	var body models.CreateCompanyModel
	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	_, err := h.Stg.User().GetById(body.OwnerId)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	id := uuid.New()

	err = h.Stg.Company().Create(id.String(), body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.Created, id)
}

func (h *Handler) GetCompanyByOwnerId(c *gin.Context) {
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

	user, err := h.Stg.User().GetById(jwtData.UserID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	company, err := h.Stg.Company().GetByOwnerId(user.ID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, company)
}
