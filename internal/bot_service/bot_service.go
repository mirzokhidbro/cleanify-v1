package bot_service

import (
	initializers "bw-erp/config"
	"bw-erp/internal/model"
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func Handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	ChatID := update.Message.Chat.ID
	var botUser = model.BotUser{ChatID: uint64(ChatID)}
	result := initializers.DB.First(&botUser)
	if update.Message.Location != nil {
		fmt.Println("heading", update.Message.Location.Heading)
		fmt.Println("HorizontalAccuracy", update.Message.Location.HorizontalAccuracy)
		fmt.Println("Latitude", update.Message.Location.Latitude)
		fmt.Println("LivePeriod", update.Message.Location.LivePeriod)
		fmt.Println("Longitude", update.Message.Location.Longitude)
		fmt.Println("ProximityAlertRadius", update.Message.Location.ProximityAlertRadius)
	}
	if result.RowsAffected == 0 || update.Message.Text == "/start" {
		RegistrationHandle(ctx, b, update)
	} else if botUser.Page == "main" {
		MainPageHandler(ctx, b, update, botUser)
	}
}
