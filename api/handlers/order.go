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

	// body.CompanyID = *user.CompanyID

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

	data, err := h.Stg.Order().GetList(body.CompanyID, models.OrdersListRequest{
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

	oldOrderStatus := order.Status

	rowsAffected, err := h.Stg.Order().Update(user.ID, &body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	order, _ = h.Stg.Order().GetByPrimaryKey(body.ID)

	// [TODO: logikani ko'rib chiqish kerak!!!]

	if body.Status != 0 && oldOrderStatus != order.Status {
		go func() {
			var wg sync.WaitGroup
			wg.Add(1)
			go func() {
				requestBody := map[string]interface{}{
					"order_id":   order.ID,
					"status":     order.Status,
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
