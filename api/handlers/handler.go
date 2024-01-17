package handlers

import (
	"bw-erp/config"
	"bw-erp/storage"
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
