package handlers

import (
	"bw-erp/api/http"
	"bw-erp/models"
	"bw-erp/pkg/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AuthUser(c *gin.Context) {
	var payload models.AuthUserModel
	if err := c.ShouldBindJSON(&payload); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	user, err := h.Stg.User().GetByPhone(payload.Phone)
	if err != nil {
		h.handleResponse(c, http.BadRequest, "Foydalanuvchi topilmadi")
		return
	}

	err = utils.VerifyPassword(user.Password, payload.Password)
	if err != nil {
		h.handleResponse(c, http.BadRequest, "Parol noto'g'ri")
		return
	}
	var response models.AuthorizationResponse

	accessToken, refreshToken, err := utils.GenerateToken(user.ID, payload.Phone)
	response.AccessToken = accessToken
	response.RefreshToken = refreshToken
	if err != err {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, response.AccessToken)
}

func (h *Handler) RefreshToken(c *gin.Context){
	
}

func (h *Handler) CurrentUser(c *gin.Context) {
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

	h.handleResponse(c, http.OK, user)
}

func (h *Handler) ChangePassword(c *gin.Context) {
	var payload models.ChangePasswordRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	jwtData, err := utils.ExtractTokenID(c)

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	user, err := h.Stg.User().GetByPhone(jwtData.Phone)

	if err != nil {
		h.handleResponse(c, http.BadRequest, user)
		return
	}

	err = utils.VerifyPassword(user.Password, payload.OldPassword)
	if err != nil {
		h.handleResponse(c, http.BadRequest, "Parol noto'g'ri!")
		return
	}
	err = h.Stg.User().ChangePassword(jwtData.UserID, payload)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, "Parol o'zgartirildi!")
}
