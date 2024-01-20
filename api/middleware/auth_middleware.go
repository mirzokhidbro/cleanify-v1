package middleware

import (
	"bw-erp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "UnAuthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
