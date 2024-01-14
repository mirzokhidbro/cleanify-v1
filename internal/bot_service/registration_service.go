package bot_service

import (
	initializers "bw-erp/config"
	"bw-erp/internal/model"
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func RegistrationHandle(ctx context.Context, b *bot.Bot, update *models.Update) {

	ChatID := update.Message.Chat.ID
	Text := ""
	var botUser = model.BotUser{ChatID: uint64(ChatID)}
	result := initializers.DB.First(&botUser)

	if result.RowsAffected != 0 {
		var botUser = model.BotUser{ChatID: uint64(ChatID)}
		initializers.DB.First(&botUser)
		botUser.FirstName = update.Message.Chat.FirstName
		botUser.LastName = update.Message.Chat.LastName
		botUser.Page = "main"
		botUser.DialogStep = "ask_section"
		initializers.DB.Save(&botUser)
	} else {
		Text = "Salom <b>" + update.Message.From.FirstName + update.Message.From.LastName + "</b> botimizga Xush kelibsiz!"
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    ChatID,
			Text:      Text,
			ParseMode: models.ParseModeHTML,
		})

		newBotUser := model.BotUser{
			ChatID:     uint64(ChatID),
			FirstName:  update.Message.Chat.FirstName,
			LastName:   update.Message.Chat.LastName,
			Page:       "main",
			DialogStep: "ask_section",
		}

		initializers.DB.Create(&newBotUser)
	}

	kb := &models.ReplyKeyboardMarkup{
		Keyboard: [][]models.KeyboardButton{
			{
				{Text: "Buyurtma"},
				{Text: "Buyurtma Elementlari"},
			},
		},
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
	}

	Text = "<b>Buyurtma</b> qismida yangi buyurtma kiritiladi \n<b>Buyurtma Elementlari</b> qismida gilamning o'lchamlari va turi kabi ma'lumotlar kiritiladi."

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      ChatID,
		Text:        Text,
		ReplyMarkup: kb,
		ParseMode:   models.ParseModeHTML,
	})
}
