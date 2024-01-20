package handlers

import (
	"bw-erp/api/http"
	"bw-erp/models"
	"bw-erp/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AuthUser(c *gin.Context) {
	var payload models.AuthUserModel
	if err := c.ShouldBindJSON(&payload); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	user, err := h.Stg.GetUserByPhone(payload.Phone)
	if err != nil {
		h.handleResponse(c, http.BadRequest, "Foydalanuvchi topilmadi")
		return
	}

	err = utils.VerifyPassword(user.Password, payload.Password)
	if err != nil {
		h.handleResponse(c, http.BadRequest, "Parol noto'g'ri")
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != err {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, token)
}

func (h *Handler) CurrentUser(c *gin.Context) {
	err := utils.TokenValid(c)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	user_id, err := utils.ExtractTokenID(c)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	user, err := h.Stg.GetUserById(user_id)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.BadRequest, user)
}
