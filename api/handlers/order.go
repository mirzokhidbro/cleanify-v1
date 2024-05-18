package handlers

import (
	"bw-erp/api/http"
	"bw-erp/models"
	"bw-erp/pkg/utils"
	"bytes"
	"encoding/json"
	newHttp "net/http"
	"strconv"
	"sync"

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

	user, err := h.Stg.User().GetById(token.UserID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	_, err = h.Stg.Company().GetById(body.CompanyID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, "company not found")
		return
	}

	body.CompanyID = *user.CompanyID

	if body.IsNewClient {
		clientID, err := h.Stg.Client().Create(models.CreateClientModel{
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

	orderID, err := h.Stg.Order().Create(user.ID, body)
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
			group, err := h.Stg.TelegramGroup().GetNotificationGroup(*user.CompanyID, 1)
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

	user, err := h.Stg.User().GetById(token.UserID)
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
	Phone := c.Query("phone")
	// ID := 0
	// if len(orderID) > 0 {
	// 	ID, err = strconv.Atoi(orderID)
	// 	if err != nil {
	// 		h.handleResponse(c, http.BadRequest, err.Error())
	// 		return
	// 	}
	// }

	data, err := h.Stg.Order().GetList(*user.CompanyID, models.OrdersListRequest{
		ID:     orderID,
		Phone:  Phone,
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
	data, err := h.Stg.Order().GetDetailedByPrimaryKey(orderId)
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

	order, err := h.Stg.Order().GetByPrimaryKey(body.ID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	rowsAffected, err := h.Stg.Order().Update(user.ID, &body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	go func() {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			requestBody := map[string]interface{}{
				"order_id":   order.ID,
				"status":     body.Status,
				"company_id": user.CompanyID,
				"flag":       h.Cfg.Release_Mode,
			}
			requestBodyJson, _ := json.Marshal(requestBody)

			url := h.Cfg.WEBHOOK_URL
			req, _ := newHttp.NewRequest("POST", url, bytes.NewBuffer(requestBodyJson))
			req.Header.Set("Content-Type", "application/json")

			client := &newHttp.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			defer wg.Done()
		}()

		wg.Wait()
	}()

	// if body.Status != 0 && order.Status != body.Status {

	// 	if body.Status != 0 && order.Status != body.Status {
	// 		BotToken := h.Cfg.BotToken
	// 		if BotToken != "" {
	// 			go func() {
	// 				opts := []bot.Option{
	// 					bot.WithDefaultHandler(h.Handler)}
	// 				group, _ := h.Stg.TelegramGroup().GetNotificationGroup(*user.CompanyID, int(body.Status))
	// 				if group.ChatID != 0 {
	// 					var Notification = ""
	// 					b, _ := bot.New(BotToken, opts...)
	// 					if err == nil {
	// 						if group.WithLocation && (order.Latitute != nil || order.Longitude != nil) && (*order.Longitude != 0 || *order.Latitute != 0) {
	// 							if body.Status == 4 || body.Status == 5 {
	// 								Notification = "Manzil: " + *order.Address + "\nTel: " + order.PhoneNumber + "\nSumma: " + strconv.FormatFloat(order.Price, 'f', -1, 64) + "\nKvadrat: " + strconv.FormatFloat(order.Square, 'f', -1, 64) + "\nIzoh: " + order.Description + "\n<a href='https://prod.yangidunyo.group/orders/" + strconv.Itoa(body.ID) + "'>Batafsil</a>"
	// 							} else {
	// 								Notification = "Manzil: " + *order.Address + "\nTel: " + order.PhoneNumber + "\nIzoh: " + order.Description + "\n<a href='https://prod.yangidunyo.group/orders/" + strconv.Itoa(body.ID) + "'>Batafsil</a>"
	// 							}
	// 							b.SendLocation(c, &bot.SendLocationParams{
	// 								ChatID:    group.ChatID,
	// 								Latitude:  *order.Latitute,
	// 								Longitude: *order.Longitude,
	// 							})
	// 							b.SendMessage(c, &bot.SendMessageParams{
	// 								ChatID:    group.ChatID,
	// 								Text:      Notification,
	// 								ParseMode: tgmodels.ParseModeHTML,
	// 							})
	// 						} else {
	// 							if body.Status == 4 || body.Status == 5 {
	// 								Notification = "Manzil: " + *order.Address + "\nTel: " + order.PhoneNumber + "\nSumma: " + strconv.FormatFloat(order.Price, 'f', -1, 64) + "\nKvadrat: " + strconv.FormatFloat(order.Square, 'f', -1, 64) + "\nIzoh: " + order.Description + "\n<a href='https://prod.yangidunyo.group/orders/" + strconv.Itoa(body.ID) + "'>Batafsil</a>"
	// 							} else {
	// 								Notification = "Manzil: " + *order.Address + "\nTel: " + order.PhoneNumber + "\nIzoh: " + order.Description + "\n<a href='https://prod.yangidunyo.group/orders/" + strconv.Itoa(body.ID) + "'>Batafsil</a>"
	// 							}
	// 							b.SendMessage(c, &bot.SendMessageParams{
	// 								ChatID:    group.ChatID,
	// 								Text:      Notification,
	// 								ParseMode: tgmodels.ParseModeHTML,
	// 							})
	// 						}
	// 					}
	// 				}
	// 			}()
	// 		}
	// 	}

	h.handleResponse(c, http.OK, rowsAffected)

	// }

}

func (h *Handler) DeleteOrder(c *gin.Context) {
	var body models.DeleteOrderRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	err := h.Stg.Order().Delete(body)

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, "Deleted successfully!")
}
