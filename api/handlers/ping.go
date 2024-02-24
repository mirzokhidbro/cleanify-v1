package handlers

import (
	"bw-erp/api/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Ping(c *gin.Context) {
	sec := time.Now().Unix()
	h.handleResponse(c, http.OK, sec)
}
