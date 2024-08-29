package handlers

import (
	"bw-erp/api/http"
	"bw-erp/models"
	"bw-erp/pkg/utils"
	"bytes"
	"encoding/json"
	"log"
	newHttp "net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
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

	if body.ClientID != 0 {
		client, _ := h.Stg.Client().GetByPrimaryKey(body.ClientID)
		if client.ID == 0 {
			h.handleResponse(c, http.BadRequest, "client not found")
			return
		}
	} else {
		client, _ := h.Stg.Client().GetByPhoneNumber(body.Phone)

		if client.ID == 0 {
			clientID, err := h.Stg.Client().Create(models.CreateClientModel{
				CompanyID:   body.CompanyID,
				PhoneNumber: body.Phone,
				Address:     body.Address,
				Longitude:   body.Longitude,
				Latitute:    body.Latitute,
			})
			if err != nil {
				h.handleResponse(c, http.BadRequest, err.Error())
				return
			}
			body.ClientID = clientID
		} else {
			body.ClientID = client.ID
		}
	}

	orderID, err := h.Stg.Order().Create(user.ID, body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	var status int8

	if body.Status == 0 {
		status = 1
	} else {
		status = body.Status
	}

	go func() {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			requestBody := map[string]interface{}{
				"order_id":   orderID,
				"status":     status,
				"company_id": body.CompanyID,
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

	h.handleResponse(c, http.Created, "Created successfully!")

}

func (h *Handler) GetOrdersList(c *gin.Context) {
	var body models.OrdersListRequest
	if err := c.ShouldBindQuery(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	data, err := h.Stg.Order().GetList(body.CompanyID, body)
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

func (h *Handler) SetOrderDiscount(c *gin.Context) {
	var body models.SetOrderPriceRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	if body.DiscountPercentage > 100 {
		h.handleResponse(c, http.BadRequest, "Invalid discount percentage")
		return
	}

	err := h.Stg.Order().SetDiscount(body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	h.handleResponse(c, http.OK, "OK!")
}

func (h *Handler) AddOrderPayment(c *gin.Context) {
	var body models.AddOrderPaymentRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	jwtData, err := utils.ExtractTokenID(c)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	user, err := h.Stg.User().GetById(jwtData.UserID)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	err = h.Stg.Order().AddPayment(user.ID, body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.OK, "OK!")
}

func (h *Handler) GetTransactionByOrder(c *gin.Context) {
	h.handleResponse(c, http.OK, "ok")
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

	oldOrderStatus := order.Status

	rowsAffected, err := h.Stg.Order().Update(user.ID, &body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	order, _ = h.Stg.Order().GetByPrimaryKey(body.ID)

	if body.Status != 0 && oldOrderStatus != order.Status {

		go func() {

			var wg sync.WaitGroup
			wg.Add(1)
			go func() {

				h.Stg.StatusChangeHistory().Create(models.CreateStatusChangeHistoryModel{
					HistoryableType: "order",
					HistoryableID:   order.ID,
					Status:          int(body.Status),
					UserID:          user.ID,
				})

				defer wg.Done()
				requestBody := map[string]interface{}{
					"order_id":   order.ID,
					"status":     order.Status,
					"company_id": order.CompanyID,
					"flag":       h.Cfg.Release_Mode,
				}
				requestBodyJson, err := json.Marshal(requestBody)
				if err != nil {
					log.Printf("Error marshalling request body: %v", err)
					return
				}

				url := h.Cfg.WEBHOOK_URL

				req, err := newHttp.NewRequest("POST", url, bytes.NewBuffer(requestBodyJson))
				if err != nil {
					log.Printf("Error creating new request: %v", err)
					return
				}
				req.Header.Set("Content-Type", "application/json")

				client := &newHttp.Client{}
				resp, err := client.Do(req)
				if err != nil {
					log.Printf("Error sending request: %v", err)
					return
				}
				defer resp.Body.Close()
			}()

			wg.Wait()
		}()
	}

	h.handleResponse(c, http.OK, rowsAffected)

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
