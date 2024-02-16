package handlers

import (
	"bw-erp/api/http"
	"bw-erp/models"
	"context"
	"os"
	"os/signal"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-telegram/bot"
	tgmodels "github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/keyboard/inline"
)

// tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// var orderKeyboard = tgbotapi.NewReplyKeyboard(
// 	tgbotapi.NewKeyboardButtonRow(
// 		tgbotapi.NewKeyboardButton("Buyurtma kiritish"),
// 	),
// )

// var locationKeyboard = tgbotapi.NewReplyKeyboard(
// 	tgbotapi.NewKeyboardButtonRow(
// 		tgbotapi.NewKeyboardButtonLocation("Lokatsiya kiritish"),
// 	),
// )

// func (h *Handler) CreateCompanyBotModel(c *gin.Context) {
// 	var body models.CreateCompanyBotModel
// 	if err := c.ShouldBindJSON(&body); err != nil {
// 		h.handleResponse(c, http.BadRequest, err.Error())
// 		return
// 	}

// 	bot, err := tgbotapi.NewBotAPI(body.BotToken)
// 	if err != nil {
// 		h.handleResponse(c, http.BadRequest, "bot token is invalid")
// 	}
// 	body.BotID = int(bot.Self.ID)
// 	body.Firstname = bot.Self.FirstName
// 	body.Lastname = bot.Self.LastName
// 	body.Username = bot.Self.UserName

// 	id := uuid.New()
// 	err = h.Stg.CreateCompanyBotModel(id.String(), body)
// 	if err != nil {
// 		h.handleResponse(c, http.BadRequest, err.Error())
// 		return
// 	}
// 	h.handleResponse(c, http.Created, id)
// 	go h.BotHandler(bot)
// }

// func (h *Handler) BotHandler(bot *tgbotapi.BotAPI) {
// 	bot.Debug = true
// 	u := tgbotapi.NewUpdate(0)
// 	u.Timeout = 10 * 60

// 	updates := bot.GetUpdatesChan(u)
// 	h.BotUpdatesHandler(updates, bot)

// }

// func (h *Handler) BotUpdatesHandler(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) {
// 	for update := range updates {
// 		if update.Message != nil {
// 			user, err := h.Stg.GetBotUserByChatIDModel(update.Message.Chat.ID, bot.Self.ID)
// 			if err != nil {
// 				msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
// 				msg.ReplyToMessageID = update.Message.MessageID
// 				bot.Send(msg)
// 				err = h.Stg.CreateBotUserModel(models.CreateBotUserModel{
// 					BotID:      int(bot.Self.ID),
// 					ChatID:     int(update.Message.Chat.ID),
// 					Page:       "Registration",
// 					DialogStep: "AskPhoneNumber",
// 				})
// 				if err != nil {
// 					msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
// 					msg.ReplyToMessageID = update.Message.MessageID
// 					bot.Send(msg)
// 				} else {
// 					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Salom "+update.Message.From.FirstName+" Iltimos tizimdagi telefon raqamingizni kiriting")
// 					msg.ReplyToMessageID = update.Message.MessageID
// 					bot.Send(msg)
// 				}
// 			} else {
// 				switch *user.Page {
// 				case "Registration":
// 					h.RegistrationPage(update, bot)
// 				case "Order Page":
// 					h.OrderPage(update, bot, user)

// 				default:
// 					fmt.Println("It's a weekday")
// 				}
// 			}
// 		}
// 	}
// }

func (h *Handler) BotStart(c *gin.Context) {
	bots, err := h.Stg.GetTelegramOrderBot()
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	go func() {
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
		defer cancel()

		var wg sync.WaitGroup

		for _, bot_config := range bots {

			opts := []bot.Option{
				bot.WithDefaultHandler(h.Handler),
			}

			b, err := bot.New(bot_config.BotToken, opts...)
			if err != nil {
				panic(err)
			}

			wg.Add(1)
			go func() {
				defer wg.Done()
				b.RegisterHandler(bot.HandlerTypeMessageText, "/olishkerak", bot.MatchTypeExact, h.newApplicationHandler)
				b.Start(ctx)
			}()
		}

		wg.Wait()
	}()

	h.handleResponse(c, http.OK, "OK!")
}

func (h *Handler) newApplicationHandler(ctx context.Context, b *bot.Bot, update *tgmodels.Update) {

	orders, err := h.Stg.GetOrdersByStatus("4a954141-5203-4e46-ae8a-519ebd098167", 0)
	if err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   err.Error(),
		})
	} else {
		kb := inline.New(b)
		for _, order := range orders {
			buttonName := *order.Address + " | " + order.Phone
			kb.Row().Button(buttonName, []byte(order.Phone), h.onInlineKeyboardSelect)
		}

		kb.Row().Button("Cancel", []byte("cancel"), h.onInlineKeyboardSelect)

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "Select the variant",
			ReplyMarkup: kb,
		})
	}

}

func (h *Handler) onInlineKeyboardSelect(ctx context.Context, b *bot.Bot, mes tgmodels.InaccessibleMessage, data []byte) {
	order, err := h.Stg.GetOrderByPhone("4a954141-5203-4e46-ae8a-519ebd098167", string(data))
	if err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: mes.Chat.ID,
			Text:   err.Error(),
		})
	} else {
		if order.Phone != "" {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: mes.Chat.ID,
				Text:   "You selected: " + order.Phone,
			})
		}

	}
}

func (h *Handler) Handler(ctx context.Context, b *bot.Bot, update *tgmodels.Update) {
	botData, _ := b.GetMe(ctx)
	botID := botData.ID
	user, err := h.Stg.GetBotUserByChatIDModel(update.Message.Chat.ID, botID)
	if err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   err.Error(),
		})
		err = h.Stg.CreateBotUserModel(models.CreateBotUserModel{
			BotID:      int(botData.ID),
			ChatID:     int(update.Message.Chat.ID),
			Page:       "Registration",
			DialogStep: "AskPhoneNumber",
		})
		if err != nil {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   err.Error(),
			})
		} else {
			text := "Salom " + update.Message.From.FirstName + " Iltimos tizimdagi telefon raqamingizni kiriting"
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   text,
			})
		}
	} else {
		switch *user.Page {
		case "Registration":
			h.RegistrationPage(ctx, b, update, botID)
		case "Order":
			h.OrderPage(ctx, b, update, user)
		}
	}
}

func (h *Handler) RegistrationPage(ctx context.Context, b *bot.Bot, update *tgmodels.Update, botID int64) {
	phone := update.Message.Text
	chatID := update.Message.Chat.ID
	user, err := h.Stg.GetSelectedUser(botID, phone)
	if err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   err.Error(),
		})
	} else {
		_, err = h.Stg.UpdateBotUserModel(models.BotUser{
			UserID: user.UserID,
			BotID:  int(botID),
			ChatID: chatID,
		})
		if err != nil {
			h.handleError(ctx, b, update, err.Error(), update.Message.MessageThreadID)
		} else {
			text := "Foydalanuvchi muvaffaqiyatli ro'yxatdan o'tkazildi"
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   text,
			})
		}
	}
}

// func (h *Handler) OrderPage(ctx context.Context, b *bot.Bot, update *tgmodels.Update, user models.BotUser) {
// 	if *user.DialogStep == "" {
// 		if update.Message.Text != "Buyurtma kiritish" {
// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Yangi buyurtma qo'shish uchun <b>Buyurtma kiritish</b> tugmasini bosing")
// 			msg.ReplyToMessageID = update.Message.MessageID
// 			bot.Send(msg)
// 		} else {
// 			h.OrderMainPage(update, bot, user)
// 		}
// 	} else {
// 		switch *user.DialogStep {
// 		case "asked order slug":
// 			h.AskedOrderSlug(update, bot, user)
// 		case "order phone number asked":
// 			h.AskedOrderPhoneNumber(update, bot, user)
// 		case "order count asked":
// 			h.AskedOrderCount(update, bot, user)
// 		case "order location asked":
// 			h.AskedOrderLocation(update, bot, user)
// 		}
// 	}
// }

// func (h *Handler) OrderMainPage(update tgbotapi.Update, bot *tgbotapi.BotAPI, user models.BotUser) {
// 	botID := bot.Self.ID
// 	chatID := update.Message.Chat.ID
// 	dialogStep := "asked order slug"
// 	_, err := h.Stg.UpdateBotUserModel(models.BotUser{
// 		UserID:     user.UserID,
// 		BotID:      int(botID),
// 		ChatID:     chatID,
// 		DialogStep: &dialogStep,
// 	})
// 	if err != nil {
// 		h.handleError(bot, chatID, err.Error(), update.Message.MessageID)
// 	} else {
// 		msg := tgbotapi.NewMessage(chatID, "Birka nomini kiriting")
// 		msg.ReplyToMessageID = update.Message.MessageID
// 		bot.Send(msg)
// 	}
// }

// func (h *Handler) AskedOrderSlug(update tgbotapi.Update, bot *tgbotapi.BotAPI, user models.BotUser) {
// 	botID := bot.Self.ID
// 	chatID := update.Message.Chat.ID
// 	dialogStep := "order phone number asked"
// 	_, err := h.Stg.UpdateBotUserModel(models.BotUser{
// 		UserID:     user.UserID,
// 		BotID:      int(botID),
// 		ChatID:     chatID,
// 		DialogStep: &dialogStep,
// 	})
// 	if err != nil {
// 		h.handleError(bot, chatID, "user "+err.Error(), update.Message.MessageID)
// 	} else {
// 		user, err := h.Stg.GetBotUserByCompany(botID, chatID)
// 		if err != nil {
// 			h.handleError(bot, chatID, "get user "+err.Error(), update.Message.MessageID)
// 		} else {
// 			id, err := h.Stg.CreateOrderModel(models.CreateOrderModel{
// 				Slug:      update.Message.Text,
// 				ChatID:    chatID,
// 				CompanyID: user.CompanyID,
// 			})
// 			if err != nil {
// 				h.handleError(bot, chatID, "create order"+err.Error(), update.Message.MessageID)
// 			} else {
// 				err = h.Stg.CreateTelegramSessionModel(models.TelegramSessionModel{
// 					BotID:   botID,
// 					ChatID:  chatID,
// 					OrderID: id,
// 				})
// 				if err != nil {
// 					h.handleError(bot, chatID, "create telegram session "+err.Error(), update.Message.MessageID)
// 				} else {
// 					msg := tgbotapi.NewMessage(chatID, "Buyurtma qabul qiluvchini telefon raqamini kiriting. \nTelefon raqam shu formatda bo'lishi kerak: +998991234567")
// 					msg.ReplyToMessageID = update.Message.MessageID
// 					bot.Send(msg)
// 				}
// 			}
// 		}

// 	}
// }

// func (h *Handler) AskedOrderPhoneNumber(update tgbotapi.Update, bot *tgbotapi.BotAPI, user models.BotUser) {
// 	botID := bot.Self.ID
// 	chatID := update.Message.Chat.ID
// 	dialogStep := "order count asked"

// 	session, err := h.Stg.GetTelegramSessionByChatIDBotID(chatID, botID)

// 	if err != nil {
// 		h.handleError(bot, chatID, err.Error(), update.Message.MessageID)
// 	} else {
// 		if !utils.IsValidPhone(update.Message.Text) {
// 			h.handleError(bot, chatID, "Telefon raqam formati noto'g'ri! Iltimos qaytadan to'g'ri formatda kiriting!", update.Message.MessageID)
// 		} else {
// 			_, err := h.Stg.UpdateBotUserModel(models.BotUser{
// 				UserID:     user.UserID,
// 				BotID:      int(botID),
// 				ChatID:     chatID,
// 				DialogStep: &dialogStep,
// 			})
// 			if err != nil {
// 				h.handleError(bot, chatID, err.Error(), update.Message.MessageID)
// 			} else {
// 				orderID := session.OrderID
// 				h.Stg.UpdateOrder(&models.UpdateOrderRequest{
// 					ID:    orderID,
// 					Phone: update.Message.Text,
// 				})
// 				msg := tgbotapi.NewMessage(chatID, "Buyurtma sonini kiriting")
// 				msg.ReplyToMessageID = update.Message.MessageID
// 				bot.Send(msg)
// 			}
// 		}
// 	}

// }

// func (h *Handler) AskedOrderCount(update tgbotapi.Update, bot *tgbotapi.BotAPI, user models.BotUser) {
// 	botID := bot.Self.ID
// 	chatID := update.Message.Chat.ID

// 	dialogStep := "order location asked"
// 	_, err := h.Stg.UpdateBotUserModel(models.BotUser{
// 		UserID:     user.UserID,
// 		BotID:      int(botID),
// 		ChatID:     chatID,
// 		DialogStep: &dialogStep,
// 	})
// 	if err != nil {
// 		h.handleError(bot, chatID, err.Error(), update.Message.MessageID)
// 	} else {
// 		session, err := h.Stg.GetTelegramSessionByChatIDBotID(chatID, botID)
// 		if err != nil {
// 			h.handleError(bot, chatID, err.Error(), update.Message.MessageID)
// 		} else {
// 			orderID := session.OrderID
// 			h.Stg.UpdateOrder(&models.UpdateOrderRequest{
// 				ID:    orderID,
// 				Count: update.Message.Text,
// 			})
// 			msg := tgbotapi.NewMessage(chatID, "Lokatsiya kiriting!")
// 			msg.ReplyToMessageID = update.Message.MessageID
// 			msg.ReplyMarkup = locationKeyboard
// 			bot.Send(msg)
// 		}
// 	}
// }

// func (h *Handler) AskedOrderLocation(update tgbotapi.Update, bot *tgbotapi.BotAPI, user models.BotUser) {
// 	botID := bot.Self.ID
// 	chatID := update.Message.Chat.ID

// 	session, err := h.Stg.GetTelegramSessionByChatIDBotID(chatID, botID)
// 	if err != nil {
// 		h.handleError(bot, chatID, err.Error(), update.Message.MessageID)
// 	} else {
// 		if update.Message.Location == nil {
// 			h.handleError(bot, chatID, "Iltimos lokatsiya kiriting!", update.Message.MessageID)
// 		} else {
// 			_, err := h.Stg.UpdateOrder(&models.UpdateOrderRequest{
// 				ID:        session.OrderID,
// 				Latitute:  update.Message.Location.Latitude,
// 				Longitude: update.Message.Location.Longitude,
// 				Status:    1,
// 			})
// 			if err != nil {
// 				h.handleError(bot, chatID, err.Error(), update.Message.MessageID)
// 			} else {
// 				dialogStep := ""
// 				_, err := h.Stg.UpdateBotUserModel(models.BotUser{
// 					UserID:     user.UserID,
// 					BotID:      int(botID),
// 					ChatID:     update.Message.Chat.ID,
// 					DialogStep: &dialogStep,
// 				})
// 				if err != nil {
// 					h.handleError(bot, chatID, err.Error(), update.Message.MessageID)
// 				} else {
// 					h.Stg.DeleteTelegramSession(session.ID)
// 					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Buyurtma qabul qilindi")
// 					msg.ReplyToMessageID = update.Message.MessageID
// 					msg.ReplyMarkup = orderKeyboard
// 					bot.Send(msg)
// 				}
// 			}
// 		}
// 	}
// }

func (h *Handler) handleError(ctx context.Context, b *bot.Bot, update *tgmodels.Update, errorMessage string, replyToMessageID int) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        errorMessage,
		ReplyMarkup: replyToMessageID,
	})
	// msg := tgbotapi.NewMessage(chatID, errorMessage)
	// msg.ReplyToMessageID = replyToMessageID
	// bot.Send(msg)
}

// func (h *Handler) SendLocation(c *gin.Context) {
// 	var query models.OrderSendLocationRequest
// 	if err := c.ShouldBindQuery(&query); err != nil {
// 		h.handleResponse(c, http.BadRequest, err.Error())
// 		return
// 	}

// 	jwtData, _ := utils.ExtractTokenID(c)
// 	order, err := h.Stg.GetOrderLocation(query.OrderID)
// 	if err != nil {
// 		h.handleResponse(c, http.OK, err.Error())
// 		return
// 	}

// 	botUser, _ := h.Stg.GetBotUserByUserID(jwtData.UserID)

// 	if order.Latitute == nil || order.Longitude == nil {
// 		h.handleResponse(c, http.OK, "Bu buyurtma lokatsiyasi mavjud emas!")
// 		return
// 	}

// 	bot, err := tgbotapi.NewBotAPI(botUser.BotToken)
// 	if err != nil {
// 		h.handleResponse(c, http.OK, err.Error())
// 		return
// 	}
// 	msg := tgbotapi.NewLocation(botUser.ChatID, *order.Latitute, *order.Longitude)
// 	bot.Send(msg)
// 	msg2 := tgbotapi.NewMessage(botUser.ChatID, "Buyurtma birkasi: "+order.Slug+"\nBuyurtmachi telefon raqami: "+order.Phone)
// 	msg2.ReplyToMessageID = msg.ReplyToMessageID
// 	bot.Send(msg2)

// 	h.handleResponse(c, http.OK, "Lokatsiya jo'natildi!")
// }
