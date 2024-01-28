package handlers

import (
	"bw-erp/api/http"
	"bw-erp/models"
	"fmt"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
)

var orderKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Buyurtma kiritish"),
	),
)

var locationKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButtonLocation("Lokatsiya kiritish"),
	),
)

func (h *Handler) CreateCompanyBotModel(c *gin.Context) {
	var body models.CreateCompanyBotModel
	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	bot, err := tgbotapi.NewBotAPI(body.BotToken)
	if err != nil {
		h.handleResponse(c, http.BadRequest, "bot token is invalid")
	}
	body.BotID = int(bot.Self.ID)

	id := uuid.New()
	err = h.Stg.CreateCompanyBotModel(id.String(), body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	h.handleResponse(c, http.Created, id)
	go h.BotHandler(bot)
}

func (h *Handler) BotHandler(bot *tgbotapi.BotAPI) {
	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 10 * 60

	updates := bot.GetUpdatesChan(u)
	h.BotUpdatesHandler(updates, bot)

}

func (h *Handler) BotUpdatesHandler(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) {
	for update := range updates {
		if update.Message != nil {
			user, err := h.Stg.GetBotUserByChatIDModel(update.Message.Chat.ID, bot.Self.ID)
			if err != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
				err = h.Stg.CreateBotUserModel(models.CreateBotUserModel{
					BotID:      int(bot.Self.ID),
					ChatID:     int(update.Message.Chat.ID),
					Page:       "Registration",
					DialogStep: "AskPhoneNumber",
				})
				if err != nil {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
					msg.ReplyToMessageID = update.Message.MessageID
					bot.Send(msg)
				} else {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Salom "+update.Message.From.FirstName+" Iltimos tizimdagi telefon raqamingizni kiriting")
					msg.ReplyToMessageID = update.Message.MessageID
					bot.Send(msg)
				}
			} else {
				switch *user.Page {
				case "Registration":
					h.RegistrationPage(update, bot)
				case "Order Page":
					h.OrderPage(update, bot, user)

				default:
					fmt.Println("It's a weekday")
				}
			}
		}
	}
}

func (h *Handler) BotStart(c *gin.Context) {
	bots, err := h.Stg.GetTelegramOrderBot()
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	for _, bot := range bots {
		newBot, err := tgbotapi.NewBotAPI(bot.BotToken)
		if err != nil {
			h.handleResponse(c, http.BadRequest, err.Error())
			return
		}
		go h.BotHandler(newBot)
	}

	h.handleResponse(c, http.OK, "OK!")
}

func (h *Handler) RegistrationPage(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	botID := bot.Self.ID
	phone := update.Message.Text
	user, err := h.Stg.GetSelectedUser(botID, phone)
	if err != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	} else {
		page := "Order Page"
		dialogStep := ""
		_, err = h.Stg.UpdateBotUserModel(models.BotUser{
			UserID:     user.UserID,
			BotID:      int(botID),
			ChatID:     update.Message.Chat.ID,
			Page:       &page,
			DialogStep: &dialogStep,
		})

		if err != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		} else {
			text := "Bu bot " + user.CompanyName + " korxonasi uchun buyurtmalarni kiritish maqsadida foydalaniladi"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			msg.ReplyToMessageID = update.Message.MessageID
			msg.ReplyMarkup = orderKeyboard
			bot.Send(msg)
		}
	}
}

func (h *Handler) OrderPage(update tgbotapi.Update, bot *tgbotapi.BotAPI, user models.BotUser) {
	if *user.DialogStep == "" {
		if update.Message.Text != "Buyurtma kiritish" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Yangi buyurtma qo'shish uchun <b>Buyurtma kiritish</b> tugmasini bosing")
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		} else {
			h.OrderMainPage(update, bot, user)
		}
	} else {
		switch *user.DialogStep {
		case "asked order slug":
			h.AskedOrderSlug(update, bot, user)
		case "order phone number asked":
			h.AskedOrderPhoneNumber(update, bot, user)
		case "order count asked":
			h.AskedOrderCount(update, bot, user)
		case "order location asked":
			h.AskedOrderLocation(update, bot, user)
		}
	}
}

func (h *Handler) OrderMainPage(update tgbotapi.Update, bot *tgbotapi.BotAPI, user models.BotUser) {
	botID := bot.Self.ID
	dialogStep := "asked order slug"
	_, err := h.Stg.UpdateBotUserModel(models.BotUser{
		UserID:     user.UserID,
		BotID:      int(botID),
		ChatID:     update.Message.Chat.ID,
		DialogStep: &dialogStep,
	})
	if err != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	} else {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Birka nomini kiriting")
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}
}

func (h *Handler) AskedOrderSlug(update tgbotapi.Update, bot *tgbotapi.BotAPI, user models.BotUser) {
	botID := bot.Self.ID
	chatID := update.Message.Chat.ID
	dialogStep := "order phone number asked"
	_, err := h.Stg.UpdateBotUserModel(models.BotUser{
		UserID:     user.UserID,
		BotID:      int(botID),
		ChatID:     chatID,
		DialogStep: &dialogStep,
	})
	if err != nil {
		msg := tgbotapi.NewMessage(chatID, "user "+ err.Error())
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	} else {
		user, err := h.Stg.GetBotUserByCompany(botID, chatID)
		if err != nil {
			msg := tgbotapi.NewMessage(chatID, "get user " + err.Error())
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		} else {
			id, err := h.Stg.CreateOrderModel(models.CreateOrderModel{
				Slug:      update.Message.Text,
				ChatID:    chatID,
				CompanyID: user.CompanyID,
			})
			if err != nil {
				msg := tgbotapi.NewMessage(chatID, "create order" + err.Error())
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			} else {
				err = h.Stg.CreateTelegramSessionModel(models.TelegramSessionModel{
					BotID:   botID,
					ChatID:  chatID,
					OrderID: id,
				})
				if err != nil {
					msg := tgbotapi.NewMessage(chatID,  "create telegram session " + err.Error())
					msg.ReplyToMessageID = update.Message.MessageID
					bot.Send(msg)
				} else {
					msg := tgbotapi.NewMessage(chatID, "Buyurtma qabul qiluvchini telefon raqamini kiriting")
					msg.ReplyToMessageID = update.Message.MessageID
					bot.Send(msg)
				}
			}
		}

	}
}

func (h *Handler) AskedOrderPhoneNumber(update tgbotapi.Update, bot *tgbotapi.BotAPI, user models.BotUser) {
	botID := bot.Self.ID
	chatID := update.Message.Chat.ID
	dialogStep := "order count asked"
	_, err := h.Stg.UpdateBotUserModel(models.BotUser{
		UserID:     user.UserID,
		BotID:      int(botID),
		ChatID:     chatID,
		DialogStep: &dialogStep,
	})
	if err != nil {
		msg := tgbotapi.NewMessage(chatID, err.Error())
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	} else {
		session, err := h.Stg.GetTelegramSessionByChatIDBotID(chatID, botID)

		if err != nil {
			msg := tgbotapi.NewMessage(chatID, err.Error())
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		} else {
			orderID := session.OrderID
			h.Stg.UpdateOrder(&models.UpdateOrderRequest{
				ID:    orderID,
				Phone: update.Message.Text,
			})
			msg := tgbotapi.NewMessage(chatID, "Buyurtma sonini kiriting")
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}

	}
}

func (h *Handler) AskedOrderCount(update tgbotapi.Update, bot *tgbotapi.BotAPI, user models.BotUser) {
	botID := bot.Self.ID
	chatID := update.Message.Chat.ID

	dialogStep := "order location asked"
	_, err := h.Stg.UpdateBotUserModel(models.BotUser{
		UserID:     user.UserID,
		BotID:      int(botID),
		ChatID:     chatID,
		DialogStep: &dialogStep,
	})
	if err != nil {
		msg := tgbotapi.NewMessage(chatID, err.Error())
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	} else {
		session, err := h.Stg.GetTelegramSessionByChatIDBotID(chatID, botID)
		if err != nil {
			msg := tgbotapi.NewMessage(chatID, "Lokatsiya kiriting!")
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		} else {
			orderID := session.OrderID
			h.Stg.UpdateOrder(&models.UpdateOrderRequest{
				ID:    orderID,
				Count: update.Message.Text,
			})
			msg := tgbotapi.NewMessage(chatID, "Lokatsiya kiriting!")
			msg.ReplyToMessageID = update.Message.MessageID
			msg.ReplyMarkup = locationKeyboard
			bot.Send(msg)
		}
	}
}

func (h *Handler) AskedOrderLocation(update tgbotapi.Update, bot *tgbotapi.BotAPI, user models.BotUser) {
	botID := bot.Self.ID
	chatID := update.Message.Chat.ID

	dialogStep := ""
	_, err := h.Stg.UpdateBotUserModel(models.BotUser{
		UserID:     user.UserID,
		BotID:      int(botID),
		ChatID:     update.Message.Chat.ID,
		DialogStep: &dialogStep,
	})
	if err != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	} else {
		session, err := h.Stg.GetTelegramSessionByChatIDBotID(chatID, botID)
		if err != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		} else {
			_, err := h.Stg.UpdateOrder(&models.UpdateOrderRequest{
				ID:        session.OrderID,
				Latitute:  update.Message.Location.Latitude,
				Longitude: update.Message.Location.Longitude,
			})
			if err != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			} else {
				h.Stg.DeleteTelegramSession(session.ID)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Buyurtma qabul qilindi")
				msg.ReplyToMessageID = update.Message.MessageID
				msg.ReplyMarkup = orderKeyboard
				bot.Send(msg)
			}
		}
	}
}
