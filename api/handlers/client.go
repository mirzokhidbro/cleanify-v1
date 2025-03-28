package handlers

import (
	"bw-erp/api/http"
	"bw-erp/models"
	"bw-erp/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-telegram/bot"
	tgmodels "github.com/go-telegram/bot/models"
)

func (h *Handler) CreateClientModel(c *gin.Context) {
	var body models.CreateClientModel
	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	if !utils.IsValidUUID(body.CompanyID) {
		h.handleResponse(c, http.InvalidArgument, "company id is an invalid uuid")
		return
	}

	_, err := h.Stg.Company().GetById(body.CompanyID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, "company not found")
		return
	}

	// token, err := utils.ExtractTokenID(c)

	// if err != nil {
	// 	h.handleResponse(c, http.BadRequest, err.Error())
	// 	return
	// }

	// user, err := h.Stg.User().GetById(token.UserID)
	// if err != nil {
	// 	h.handleResponse(c, http.BadRequest, err.Error())
	// 	return
	// }

	// body.CompanyID = *user.CompanyID

	_, err = h.Stg.Client().Create(body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.Created, "Created successfully!")
}

func (h *Handler) GetClientsList(c *gin.Context) {
	var body models.ClientListRequest
	if err := c.ShouldBindQuery(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	// token, err := utils.ExtractTokenID(c)

	// if err != nil {
	// 	h.handleResponse(c, http.BadRequest, err.Error())
	// 	return
	// }

	// user, err := h.Stg.User().GetById(token.UserID)
	// if err != nil {
	// 	h.handleResponse(c, http.BadRequest, err.Error())
	// 	return
	// }

	data, err := h.Stg.Client().GetList(body.CompanyID, body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, data)
}

func (h *Handler) GetClientByPrimaryKey(c *gin.Context) {
	clientID := c.Param("client-id")
	clientId, err := strconv.Atoi(clientID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	data, err := h.Stg.Client().GetByPrimaryKey(clientId)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	h.handleResponse(c, http.OK, data)
}

func (h *Handler) UpdateClient(c *gin.Context) {
	var body models.UpdateClientRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	if !utils.IsValidUUID(body.CompanyID) {
		h.handleResponse(c, http.InvalidArgument, "company id is an invalid uuid")
		return
	}

	_, err := h.Stg.Company().GetById(body.CompanyID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, "company not found")
		return
	}

	// token, err := utils.ExtractTokenID(c)

	// if err != nil {
	// 	h.handleResponse(c, http.BadRequest, err.Error())
	// 	return
	// }

	// user, err := h.Stg.User().GetById(token.UserID)
	// if err != nil {
	// 	h.handleResponse(c, http.BadRequest, err.Error())
	// 	return
	// }

	// body.CompanyID = *user.CompanyID

	_, err = h.Stg.Client().Update(body.CompanyID, body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.Created, "Edited successfully!")
}

func (h *Handler) SetLocation(c *gin.Context) {
	clientID := c.Param("client-id")
	clientId, err := strconv.Atoi(clientID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	data, err := h.Stg.Client().GetByPrimaryKey(clientId)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	token, err := utils.ExtractTokenID(c)

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	user, err := h.Stg.User().GetById(token.UserID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	botUser, err := h.Stg.BotUser().GetByUserID(user.ID)

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	err = h.Stg.TelegramSession().Create(models.TelegramSessionModel{
		BotID:   int64(botUser.BotID),
		ChatID:  botUser.ChatID,
		OrderID: data.ID,
	})
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	opts := []bot.Option{
		bot.WithDefaultHandler(h.Handler),
	}

	page := "SetLocation"
	h.Stg.BotUser().Update(models.BotUser{
		UserID: &user.ID,
		BotID:  int(botUser.BotID),
		ChatID: botUser.ChatID,
		Page:   &page,
	})

	b, _ := bot.New(botUser.BotToken, opts...)
	Notification := "Manzil: " + data.Address + "\nTelefon raqam: " + data.PhoneNumber + "\nLokatsiya kiriting"
	kb := &tgmodels.ReplyKeyboardMarkup{
		Keyboard: [][]tgmodels.KeyboardButton{
			{
				{Text: "Lokatsiya", RequestLocation: true},
			},
		},
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
	}
	b.SendMessage(c, &bot.SendMessageParams{
		ChatID:      botUser.ChatID,
		Text:        Notification,
		ReplyMarkup: kb,
	})

	h.handleResponse(c, http.OK, "OK!")
}
