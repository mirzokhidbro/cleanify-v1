package handlers

import (
	"bw-erp/api/http"
	"bw-erp/pkg/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetWorkVolumeList(c *gin.Context) {
	// companyID := c.Param("company-id")
	// if !utils.IsValidUUID(companyID) {
	// 	h.handleResponse(c, http.InvalidArgument, "company id is an invalid uuid")
	// 	return
	// }

	token, err := utils.ExtractTokenID(c)

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	user, err := h.Stg.GetUserById(token.UserID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	data, err := h.Stg.GetWorkVolumeList(*user.CompanyID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, data)
}
