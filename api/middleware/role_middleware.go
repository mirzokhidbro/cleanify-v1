package middleware

import (
	"bw-erp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(roles string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// var sep string = ","
		// rolesList := strings.Split(roles, sep)
		err := utils.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}

		// user_id, err := utils.ExtractTokenID(c)
		// if err != nil {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		// 	c.Abort()
		// 	return
		// }
		// exist, _ := helper.InArray(currentUser.Role, rolesList)

		// if !exist {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"message": "You have not a right permission"})
		// 	c.Abort()
		// 	return
		// }

		c.Next()

	}
}
