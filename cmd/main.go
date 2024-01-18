package main

import (
	"bw-erp/api"
	"bw-erp/api/handlers"
	"bw-erp/config"
	"bw-erp/storage"
	"bw-erp/storage/postgres"
	"fmt"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig()

	psqlConnString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUserName,
		cfg.DBUserPassword,
		cfg.DBName,
	)

	fmt.Printf("variable port=%v is of type %T \n", cfg.DBPort, cfg.DBPort)

	var stg storage.StorageI
	stg, err = postgres.InitDB(psqlConnString)
	if err != nil {
		panic(err)
	}

	h := handlers.NewHandler(stg, cfg)

	r := api.SetUpRouter(h, cfg)
	server := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: r,
	}
	server.ListenAndServe()
}
