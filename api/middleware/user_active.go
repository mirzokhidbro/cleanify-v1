package middleware

import (
	"bw-erp/pkg/utils"
	"bw-erp/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserActiveMiddleware(stg storage.StorageI) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.TokenValid(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code":    http.StatusForbidden,
				"message": err.Error(),
			})
			return
		}

		jwtData, err := utils.ExtractTokenID(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}

		user, err := stg.User().GetById(jwtData.UserID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}

		if !user.IsActive {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusServiceUnavailable,
				"message": "user has been deactivated",
			})
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
