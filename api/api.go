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
		usersRouter.POST("", h.CreateUser)
		usersRouter.GET("", h.GetUsersList)
	}

	{
		authRouter := baseRouter.Group("/auth")
		authRouter.POST("/login", h.AuthUser)
		authRouter.POST("/me", h.CurrentUser)
	}

	{
		companyRouter := baseRouter.Group("/company")
		companyRouter.POST("", h.CreateCompanyModel)
		companyRouter.Use(middleware.AuthMiddleware()).GET("/get-by-owner", h.GetCompanyByOwnerId)
	}

	{
		companyRole := baseRouter.Group("/company-role")
		companyRole.POST("", h.CreateCompanyRoleModel)
		companyRole.Use(middleware.AuthMiddleware()).GET("/:company-id", h.GetRolesListByCompany)
	}

	return
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
