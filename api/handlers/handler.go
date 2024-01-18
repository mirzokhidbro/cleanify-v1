package handlers

import (
	"bw-erp/api/http"
	"bw-erp/config"
	"bw-erp/storage"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Stg storage.StorageI
	Cfg config.Config
}

func NewHandler(stg storage.StorageI, cfg config.Config) Handler {
	return Handler{
		Stg: stg,
		Cfg: cfg,
	}
}

func (h *Handler) handleResponse(c *gin.Context, status http.Status, data interface{}) {
	c.JSON(status.Code, http.Response{
		Status:      status.Status,
		Description: status.Description,
		Data:        data,
	})
}
