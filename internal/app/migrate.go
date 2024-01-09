package app

import (
	"fmt"
	"log"

	"bw-erp/config"
)

func RunMigration() {
	config, err := initializers.LoadConfig()

	if err != nil {
		log.Fatal(err.Error())
	}

	initializers.ConnectDB(&config)

	initializers.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	// config.DB.AutoMigrate(&model.User{}, &model.Company{}, &model.Employee{})
	fmt.Println("👍 Migration complete")
}
