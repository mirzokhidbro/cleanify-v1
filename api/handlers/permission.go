package handlers

import (
	"bw-erp/api/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPermissionList(c *gin.Context) {
	// token, err := utils.ExtractTokenID(c)

	// if err != nil {
	// 	h.handleResponse(c, http.BadRequest, err.Error())
	// 	return
	// }

	// user, err := h.Stg.User().GetById(token.UserID)

	// if err != nil {
	// 	h.handleResponse(c, http.BadRequest, err.Error())
	// 	return
	// }

	data, err := h.Stg.Permission().GetList("company")
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, data)
}
