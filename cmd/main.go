package main

import (
	"bw-erp/api"
	"bw-erp/api/handlers"
	"bw-erp/config"
	"bw-erp/storage"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

func main() {
	cfg, _ := config.LoadConfig()

	psqlConnString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUserName,
		cfg.DBUserPassword,
		cfg.DBName,
	)

	var stg storage.StorageI
	tempDB, err := sqlx.Connect("postgres", psqlConnString)
	if err != nil {
		panic(err)
	}
	stg = storage.NewStoragePg(tempDB)

	h := handlers.NewHandler(stg, cfg)

	r := api.SetUpRouter(h, cfg)
	server := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: r,
	}
	server.ListenAndServe()
}
