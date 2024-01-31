package middleware

import (
	"bw-erp/api/http"
	"bw-erp/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.TokenValid(c)
		if err != nil {
			c.JSON(http.Forbidden.Code, http.Response{
				Status:      "Forbidden",
				Description: "Unauthorized",
				Data:        nil,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
