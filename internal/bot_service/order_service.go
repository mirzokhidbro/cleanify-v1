package bot_service

import (
	initializers "bw-erp/config"
	"bw-erp/internal/model"
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func MainPageHandler(ctx context.Context, b *bot.Bot, update *models.Update, botUser model.BotUser) {
	received_message := ""
	sending_message := ""
	if botUser.DialogStep == "ask_section" {
		received_message = update.Message.Text
		fmt.Println(received_message)
		if received_message == "Buyurtma" {
			botUser.DialogStep = "order_slug_asked"
			initializers.DB.Save(&botUser)
			sending_message = "Buyurtmani birka yozuvini kiriting"
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:    botUser.ChatID,
				Text:      sending_message,
				ParseMode: models.ParseModeHTML,
			})
		} else if received_message == "Buyurtma Elementlari" {
			botUser.DialogStep = "order_id_asked"
			initializers.DB.Save(&botUser)
			sending_message = "lokatsiya kiriting"
			kb := &models.ReplyKeyboardMarkup{
				Keyboard: [][]models.KeyboardButton{
					{
						{Text: "Lokatsiya kiriting", RequestLocation: true},
					},
				},
				ResizeKeyboard:  true,
				OneTimeKeyboard: true,
			}
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:      botUser.ChatID,
				Text:        sending_message,
				ReplyMarkup: kb,
			})
		}
	} else if botUser.DialogStep == "order_slug_asked" {
		botUser.DialogStep = "order_phone_asked"
		initializers.DB.Save(&botUser)
		sending_message = "Buyurtma qabul qiluvchining telefon raqamini kiriting \nRaqam mana shu formatda bo'lishi kerak <b>+998991234567</b>"
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: botUser.ChatID,
			Text:   sending_message,
		})
	} else if botUser.DialogStep == "order_phone_asked" {
		botUser.DialogStep = "order_count_asked"
		initializers.DB.Save(&botUser)
		sending_message = "Buyurtma sonini kiriting"
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: botUser.ChatID,
			Text:   sending_message,
		})
	} else if botUser.DialogStep == "order_count_asked" {
		botUser.DialogStep = "ask_section"
		botUser.Page = "main"
		initializers.DB.Save(&botUser)
		sending_message = "Buyurtma muvaffaqiyatli saqlandi"
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: botUser.ChatID,
			Text:   sending_message,
		})
	}
}
