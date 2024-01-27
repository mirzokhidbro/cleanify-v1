package handlers

import (
	"bw-erp/api/http"
	"bw-erp/models"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
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
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	h.BotUpdatesHandler(updates, bot)

}

func (h *Handler) BotUpdatesHandler(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) {
	for update := range updates {
		if update.Message != nil {
			user, err := h.Stg.GetBotUserByChatIDModel(update.Message.Chat.ID)
			if err != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()+"||"+*user.DialogStep)
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
				err = h.Stg.CreateBotUserModel(models.CreateBotUserModel{
					BotID:      int(bot.Self.ID),
					ChatID:     int(update.Message.Chat.ID),
					Page:       "Registration",
					DialogStep: "AskUsername",
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
			}

			if *user.Page == "Registration" {
				h.RegistrationPage(update, bot)
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
		_, err = h.Stg.UpdateBotUserModel(models.BotUser{
			UserID: user.UserID,
			BotID:  int(botID),
		})

		if err != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		} else {
			// text := "bu bot"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Foydalanuvchi tasdiqlandi!")
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}

	}

}
