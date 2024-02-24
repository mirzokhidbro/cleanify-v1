package handlers

import (
	"bw-erp/api/http"
	"bw-erp/models"
	"bw-erp/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-telegram/bot"
)

func (h *Handler) CreateOrderModel(c *gin.Context) {
	var body models.CreateOrderModel
	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	token, err := utils.ExtractTokenID(c)

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	user, err := h.Stg.GetUserById(token.UserID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	body.CompanyID = *user.CompanyID

	_, err = h.Stg.CreateOrderModel(body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	botUser, _ := h.Stg.GetBotUserByUserID(user.ID)
	opts := []bot.Option{
		bot.WithDefaultHandler(h.Handler),
	}
	group, err := h.Stg.GetNotificationGroup(*user.CompanyID)

	if err == nil {
		b, _ := bot.New(botUser.BotToken, opts...)
		Notification := "#zayavka\nManzil: " + body.Address + "\nTel: " + body.Phone
		b.SendMessage(c, &bot.SendMessageParams{
			ChatID: group.ChatID,
			Text:   Notification,
		})
	}

	h.handleResponse(c, http.Created, "Created successfully!")

}

func (h *Handler) GetOrdersList(c *gin.Context) {
	// companyID := c.Param("company-id")
	// if !utils.IsValidUUID(companyID) {
	// 	h.handleResponse(c, http.InvalidArgument, "company id is an invalid uuid")
	// 	return
	// }
	limit, err := h.getLimitParam(c)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	offset, err := h.getOffsetParam(c)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	token, err := utils.ExtractTokenID(c)

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	user, err := h.Stg.GetUserById(token.UserID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	status, err := h.getStatusParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}
	data, err := h.Stg.GetOrdersList(*user.CompanyID, models.OrdersListRequest{
		Slug:   c.Query("slug"),
		Status: status,
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, data)
}

func (h *Handler) GetOrderByPrimaryKey(c *gin.Context) {
	orderID := c.Param("order-id")
	orderId, err := strconv.Atoi(orderID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	data, err := h.Stg.GetOrderByPrimaryKey(orderId)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	h.handleResponse(c, http.OK, data)
}

func (h *Handler) UpdateOrderModel(c *gin.Context) {
	var body models.UpdateOrderRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	rowsAffected, err := h.Stg.UpdateOrder(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, rowsAffected)

}
