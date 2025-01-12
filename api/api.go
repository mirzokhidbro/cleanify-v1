package api

import (
	"bw-erp/api/handlers"
	"bw-erp/api/middleware"
	"bw-erp/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(h handlers.Handler, cfg config.Config) (r *gin.Engine) {
	r = gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.Use(customCORSMiddleware())

	r.GET("api/ping", h.Ping)
	r.GET("api/get-order-receipt-by-uuid/:uuid", h.GetOrderReceiptByUuid)

	baseRouter := r.Group("/api/v1")
	{
		{
			authRouter := baseRouter.Group("/auth")
			authRouter.POST("/login", h.AuthUser)
			authRouter.POST("/me", h.CurrentUser)
			authRouter.POST("/refresh-token", h.RefreshToken)
			authRouter.Use(middleware.AuthMiddleware()).POST("/change-password", h.ChangePassword)
		}

		notificationRouter := baseRouter.Group("/notifications")
		{
			notificationRouter.GET("/ws", h.HandleNotificationWebSocket) // WebSocket endpoint

			// Web Push endpoints
			webpushRouter := notificationRouter.Group("/webpush")
			webpushRouter.Use(middleware.AuthMiddleware())
			{
				webpushRouter.POST("/subscribe", h.SavePushSubscription)
				// webpushRouter.GET("/subscription", h.GetPushSubscription)
				// webpushRouter.DELETE("/subscription", h.DeletePushSubscription)
				// webpushRouter.GET("/subscriptions", h.GetAllPushSubscriptions)
				// webpushRouter.POST("/send", h.SendPushNotification)
			}

			notificationRouter.Use(middleware.AuthMiddleware()).GET("list", h.GetMyNotifications)
			notificationRouter.Use(middleware.AuthMiddleware()).GET("/unread-notifications-count", h.UnreadNotificationsCount)
		}

		usersRouter := baseRouter.Group("/users")
		usersRouter.Use(middleware.AuthMiddleware()).POST("", h.Create)
		usersRouter.Use(middleware.AuthMiddleware()).GET("", h.GetList)
		usersRouter.Use(middleware.AuthMiddleware()).POST("/edit", h.Edit)
		usersRouter.Use(middleware.AuthMiddleware()).GET("/:user-id", h.GetById)
		usersRouter.Use(middleware.AuthMiddleware()).GET("/couriers", h.GetCouriesList)

		usersRouter.Use(middleware.AuthMiddleware()).POST("/employees", h.CreateEmployee)
		usersRouter.Use(middleware.AuthMiddleware()).GET("/employees", h.GetEmployeeList)
		usersRouter.Use(middleware.AuthMiddleware()).GET("/employees/show", h.ShowEmployeeDetailedData)
		usersRouter.Use(middleware.AuthMiddleware()).POST("/employees/add-transaction", h.AddTransaction)
		usersRouter.Use(middleware.AuthMiddleware()).POST("/employees/attendance", h.Attendance)
	}

	baseRouter.Static("/uploads", "./uploads")

	{
		companyRouter := baseRouter.Group("/company")
		companyRouter.POST("", h.CreateCompanyModel)
	}

	{
		orderRouter := baseRouter.Group("orders")
		orderRouter.Use(middleware.AuthMiddleware()).POST("", h.CreateOrderModel)
		orderRouter.Use(middleware.AuthMiddleware()).GET("", h.GetOrdersList)
		orderRouter.Use(middleware.AuthMiddleware()).GET("/:order-id", h.GetOrderByPrimaryKey)
		orderRouter.Use(middleware.AuthMiddleware()).POST("/edit", h.UpdateOrderModel)
		orderRouter.Use(middleware.AuthMiddleware()).POST("/set-price", h.SetOrderPrice)
		orderRouter.Use(middleware.AuthMiddleware()).POST("add-payment", h.AddOrderPayment)
		orderRouter.Use(middleware.AuthMiddleware()).GET("get-transactions-by-order", h.GetTransactionByOrder)
		orderRouter.Use(middleware.AuthMiddleware()).DELETE("", h.DeleteOrder)
		orderRouter.Use(middleware.AuthMiddleware()).POST("/comment", h.AddOrderComment)
	}

	{
		notificationSettingRouter := baseRouter.Group("notification-setting")
		notificationSettingRouter.Use(middleware.AuthMiddleware()).POST("", h.SetNotificationSetting)
		notificationSettingRouter.Use(middleware.AuthMiddleware()).GET("", h.UsersListForNotificationSettings)
		notificationSettingRouter.Use(middleware.AuthMiddleware()).GET("get-users-by-status", h.GetUsersByStatus)
	}

	{
		orderStatuses := baseRouter.Group("order-statuses")
		orderStatuses.Use(middleware.AuthMiddleware()).GET("", h.GetOrderStatusesList) //
		orderStatuses.Use(middleware.AuthMiddleware()).PUT("", h.UpdateOrderStatusModel)
		orderStatuses.Use(middleware.AuthMiddleware()).GET("/get-by-primary-key/:id", h.GetOrderStatusById)
	}

	{
		orderItemRouter := baseRouter.Group("order-items")
		orderItemRouter.Use(middleware.AuthMiddleware()).POST("", h.CreateOrderItemModel)
		orderItemRouter.Use(middleware.AuthMiddleware()).POST("edit", h.UpdateOrderItemModel)
		orderItemRouter.Use(middleware.AuthMiddleware()).DELETE("/:id", h.DeleteOrderItemByID)
		orderItemRouter.Use(middleware.AuthMiddleware()).POST("/edit-status", h.UpdateOrderItemStatus)
	}

	{
		orderItemTypeRouter := baseRouter.Group("order-item-type")
		orderItemTypeRouter.Use(middleware.AuthMiddleware()).POST("", h.CreateOrderItemTypeModel)
		orderItemTypeRouter.Use(middleware.AuthMiddleware()).GET("", h.GetOrderItemTypesByCompany) //
		orderItemTypeRouter.Use(middleware.AuthMiddleware()).PUT("", h.UpdateOrderItemType)
		orderItemTypeRouter.Use(middleware.AuthMiddleware()).GET("get-by-primary-key/:id", h.GetOrderItemTypeByID)
	}

	{
		statistics := baseRouter.Group("statistics")
		statistics.Use(middleware.AuthMiddleware()).GET("work-volume", h.GetWorkVolumeList) //
		statistics.Use(middleware.AuthMiddleware()).GET("get-service-statistics-payment", h.GetServicePaymentStatistics)
	}
	{
		statistics := baseRouter.Group("permissions")
		statistics.Use(middleware.AuthMiddleware()).GET("", h.GetPermissionList)
	}

	{
		clientRouter := baseRouter.Group("/client")
		clientRouter.Use(middleware.AuthMiddleware()).POST("", h.CreateClientModel) //
		clientRouter.Use(middleware.AuthMiddleware()).GET("/get-by-primary-key/:client-id", h.GetClientByPrimaryKey)
		clientRouter.Use(middleware.AuthMiddleware()).GET("", h.GetClientsList) //
		clientRouter.Use(middleware.AuthMiddleware()).GET("/set-location/:client-id", h.SetLocation)
		clientRouter.Use(middleware.AuthMiddleware()).PUT("", h.UpdateClient) //
	}

	{
		telegramGroup := baseRouter.Group("/telegram-group")
		telegramGroup.Use(middleware.AuthMiddleware()).POST("/verification", h.VerificationGroup) //
		telegramGroup.Use(middleware.AuthMiddleware()).GET("", h.GetTelegramGroupList)            //
		telegramGroup.Use(middleware.AuthMiddleware()).GET("/get-by-primary-key/:id", h.GetTelegramGroupByPrimaryKey)
		telegramGroup.Use(middleware.AuthMiddleware()).PUT("/:id", h.UpdateTelegramGroup)
	}

	return
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		// Add WebSocket specific headers
		if c.Request.Header.Get("Upgrade") == "websocket" {
			c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With, Connection, Upgrade, Sec-WebSocket-Key, Sec-WebSocket-Version, Sec-WebSocket-Extensions")
		} else {
			c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
		}

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
