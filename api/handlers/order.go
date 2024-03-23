package handlers

import (
	"bw-erp/api/http"
	"bw-erp/models"
	"bw-erp/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-telegram/bot"
	tgmodels "github.com/go-telegram/bot/models"
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

	if body.IsNewClient {
		clientID, err := h.Stg.CreateClientModel(models.CreateClientModel{
			CompanyID:   body.CompanyID,
			PhoneNumber: body.Phone,
			Address:     body.Address,
			Longitude:   body.Longitude,
			Latitute:    body.Latitute,
		})
		body.ClientID = clientID
		if err != nil {
			h.handleResponse(c, http.BadRequest, err.Error())
			return
		}
	} else if body.ClientID == 0 {
		h.handleResponse(c, http.BadRequest, "client id is required")
		return
	}

	orderID, err := h.Stg.CreateOrderModel(body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	BotToken := h.Cfg.BotToken
	if BotToken != "" {
		go func() {
			opts := []bot.Option{
				bot.WithDefaultHandler(h.Handler),
			}
			group, err := h.Stg.GetNotificationGroup(*user.CompanyID, 83)
			if err == nil {
				b, _ := bot.New(BotToken, opts...)
				Notification := "#zayavka\nManzil: " + body.Address + "\nTel: " + body.Phone + "\nIzoh:" + body.Description + "\n<a href='https://prod.yangidunyo.group/orders/" + strconv.Itoa(orderID) + "'>Batafsil</a>"
				b.SendMessage(c, &bot.SendMessageParams{
					ChatID:    group.ChatID,
					Text:      Notification,
					ParseMode: tgmodels.ParseModeHTML,
				})
			}
		}()
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
	orderID := c.Query("id")
	ID := 0
	if len(orderID) > 0 {
		ID, err = strconv.Atoi(orderID)
		if err != nil {
			h.handleResponse(c, http.BadRequest, err.Error())
			return
		}
	}
	var (
		DateFrom time.Time
		DateTo   time.Time
	)

	DateFromQuery := c.Query("date_from")
	if DateFromQuery != "" {
		DateFrom, err = utils.StringToTime(DateFromQuery)
		if err != nil {
			h.handleResponse(c, http.BadRequest, err.Error())
			return
		}
	}

	DateToQuery := c.Query("date_to")
	if DateToQuery != "" {
		DateTo, err = utils.StringToTime(DateToQuery)
		if err != nil {
			h.handleResponse(c, http.BadRequest, err.Error())
			return
		}
	}

	data, err := h.Stg.GetOrdersList(*user.CompanyID, models.OrdersListRequest{
		ID:       ID,
		Status:   status,
		Limit:    int32(limit),
		Offset:   int32(offset),
		DateFrom: DateFrom,
		DateTo:   DateTo,
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
	data, err := h.Stg.GetOrderDetailedByPrimaryKey(orderId)
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

	if body.Status != 0 {
		BotToken := h.Cfg.BotToken
		if BotToken != "" {
			go func() {
				opts := []bot.Option{
					bot.WithDefaultHandler(h.Handler),
				}
				group, _ := h.Stg.GetNotificationGroup(*user.CompanyID, int(body.Status))
				if group.ChatID != 0 {
					order, err := h.Stg.GetOrderByPrimaryKey(body.ID)
					b, _ := bot.New(BotToken, opts...)
					if err == nil {
						if group.WithLocation && (order.Latitute != nil || order.Longitude != nil) {
							b.SendLocation(c, &bot.SendLocationParams{
								ChatID:    group.ChatID,
								Latitude:  *order.Latitute,
								Longitude: *order.Longitude,
							})
							Notification := "Manzil: " + *order.Address + "\nTel: " + order.PhoneNumber + "\n<a href='https://prod.yangidunyo.group/orders/" + strconv.Itoa(body.ID) + "'>Batafsil</a>"
							b.SendMessage(c, &bot.SendMessageParams{
								ChatID:    group.ChatID,
								Text:      Notification,
								ParseMode: tgmodels.ParseModeHTML,
							})
						} else {
							Notification := "Manzil: " + *order.Address + "\nTel: " + order.PhoneNumber + "\n<a href='https://prod.yangidunyo.group/orders/" + strconv.Itoa(body.ID) + "'>Batafsil</a>"
							b.SendMessage(c, &bot.SendMessageParams{
								ChatID:    group.ChatID,
								Text:      Notification,
								ParseMode: tgmodels.ParseModeHTML,
							})
						}
					}
				}
			}()
		}
	}

	h.handleResponse(c, http.OK, rowsAffected)

}
