package handlers

import (
	"bw-erp/api/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Ping(c *gin.Context) {
	h.handleResponse(c, http.OK, "Ping!!!!!!")
}
