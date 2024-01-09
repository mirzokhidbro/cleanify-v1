package main

import (
	"bw-erp/helper"
	"bw-erp/internal/app"
	"net/http"
)

func main() {

	app.RunMigration()

	routes := app.NewRouter()

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)

}
