package api

import (
	"bw-erp/api/handlers"
	"bw-erp/api/middleware"
	"bw-erp/config"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(h handlers.Handler, cfg config.Config) (r *gin.Engine) {
	r = gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.Use(customCORSMiddleware())

	r.GET("api/ping", h.Ping)

	baseRouter := r.Group("/api/v1")
	{
		usersRouter := baseRouter.Group("/users")
		usersRouter.POST("", h.Create)
		usersRouter.Use(middleware.AuthMiddleware()).GET("", h.GetList)
		usersRouter.Use(middleware.AuthMiddleware()).PUT("/", h.Edit) //
		usersRouter.Use(middleware.AuthMiddleware()).GET("/:user-id", h.GetById)
	}

	{
		authRouter := baseRouter.Group("/auth")
		authRouter.POST("/login", h.AuthUser)
		authRouter.POST("/me", h.CurrentUser)
		authRouter.Use(middleware.AuthMiddleware()).POST("/change-password", h.ChangePassword)
	}

	{
		companyRouter := baseRouter.Group("/company")
		companyRouter.POST("", h.CreateCompanyModel)
		// companyRouter.Use(middleware.AuthMiddleware()).GET("/get-by-owner", h.GetCompanyByOwnerId)
	}

	{
		roleRouter := baseRouter.Group("/role")
		roleRouter.POST("", h.CreateRoleModel)
		roleRouter.Use(middleware.AuthMiddleware()).GET("/show/:role-id", h.GetRoleByPrimaryKey)
		roleRouter.Use(middleware.AuthMiddleware()).GET("/:company-id", h.GetRolesListByCompany)
		roleRouter.Use(middleware.AuthMiddleware()).POST("/give-permissions", h.GetPermissionsToRole)
	}

	{
		orderRouter := baseRouter.Group("orders")
		orderRouter.Use(middleware.AuthMiddleware()).POST("", h.CreateOrderModel)
		orderRouter.Use(middleware.AuthMiddleware()).GET("", h.GetOrdersList)
		orderRouter.Use(middleware.AuthMiddleware()).GET("/:order-id", h.GetOrderByPrimaryKey)
		orderRouter.Use(middleware.AuthMiddleware()).POST("/edit", h.UpdateOrderModel)
		orderRouter.Use(middleware.AuthMiddleware()).GET("/send-location", h.SendLocation)
		orderRouter.Use(middleware.AuthMiddleware()).DELETE("", h.DeleteOrder)
	}

	{
		orderStatuses := baseRouter.Group("order-statuses")
		orderStatuses.Use(middleware.AuthMiddleware()).GET("/:company-id", h.GetOrderStatusesList)
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
		orderItemTypeRouter.Use(middleware.AuthMiddleware()).GET("/:company-id", h.GetOrderItemTypesByCompany)
		orderItemTypeRouter.Use(middleware.AuthMiddleware()).PUT("", h.UpdateOrderItemType)
		orderItemTypeRouter.Use(middleware.AuthMiddleware()).GET("get-by-primary-key/:id", h.GetOrderItemTypeByID)
	}

	{
		companyBotRouter := baseRouter.Group("company-bot")
		companyBotRouter.Use(middleware.AuthMiddleware()).POST("", h.CreateCompanyBotModel)
		companyBotRouter.Use(middleware.AuthMiddleware()).GET("/start", h.BotStart)
	}
	{
		statistics := baseRouter.Group("statistics")
		statistics.Use(middleware.AuthMiddleware()).GET("work-volume/:company-id", h.GetWorkVolumeList) //
	}
	{
		statistics := baseRouter.Group("permissions")
		statistics.Use(middleware.AuthMiddleware()).GET("", h.GetPermissionList)
	}

	{
		clientRouter := baseRouter.Group("/client")
		clientRouter.Use(middleware.AuthMiddleware()).POST("/:company-id", h.CreateClientModel)
		clientRouter.Use(middleware.AuthMiddleware()).GET("/get-by-primary-key/:client-id", h.GetClientByPrimaryKey)
		clientRouter.Use(middleware.AuthMiddleware()).GET("/:company-id", h.GetClientsList)
		clientRouter.Use(middleware.AuthMiddleware()).GET("/set-location/:client-id", h.SetLocation)
		clientRouter.Use(middleware.AuthMiddleware()).PUT("/:company-id", h.UpdateClient)
	}

	{
		telegramGroup := baseRouter.Group("/telegram-group")
		telegramGroup.Use(middleware.AuthMiddleware()).POST("/verification/:company-id", h.VerificationGroup)
		telegramGroup.Use(middleware.AuthMiddleware()).GET("/:company-id", h.GetTelegramGroupList)
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
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, Origin, Cache-Control, X-Requested-With,  Platform-Type")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
