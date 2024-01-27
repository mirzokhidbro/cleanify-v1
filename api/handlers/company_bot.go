package handlers

import (
	"bw-erp/api/http"
	"bw-erp/models"
	"log"

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

	log.Printf("Authorized on account %s", bot.Self.ID)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "salom")
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
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
