package app

import (
	"fmt"
	"log"

	initializers "bw-erp/config"
	"bw-erp/internal/model"
)

func RunMigration() {
	config, err := initializers.LoadConfig()

	if err != nil {
		log.Fatal(err.Error())
	}

	initializers.ConnectDB(&config)

	initializers.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	initializers.DB.AutoMigrate(model.BotUser{}, model.Order{})
	fmt.Println("👍 Migration complete")
}
