package handlers

import (
	"bw-erp/api/http"
	"bw-erp/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetWorkVolumeList(c *gin.Context) {
	companyID := c.Param("company-id")
	if !utils.IsValidUUID(companyID) {
		h.handleResponse(c, http.InvalidArgument, "company id is an invalid uuid")
		return
	}

	data, err := h.Stg.GetWorkVolumeList(companyID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, data)
}
