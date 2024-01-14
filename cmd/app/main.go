package main

import (
	"bw-erp/helper"
	"bw-erp/internal/app"
	"net/http"
	"os"
)

func main() {

	app.RunMigration()

	routes := app.NewRouter()
	port := os.Getenv("PORT")

	server := &http.Server{
		Addr:    ":" + port,
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)

}
