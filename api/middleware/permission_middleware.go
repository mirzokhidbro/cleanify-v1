package middleware

// func PermissionMiddleware(h handlers.Handler, permission string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		jwtData, err := utils.ExtractTokenID(c)
// 		if err != nil {
// 			c.JSON(http.Forbidden.Code, http.Response{
// 				Status:      "Forbidden",
// 				Description: err.Error(),
// 				Data:        nil,
// 			})
// 			c.Abort()
// 			return
// 		}

// 		user, err := h.Stg.User().GetById(jwtData.UserID)
// 		if err != nil {
// 			c.JSON(http.Forbidden.Code, http.Response{
// 				Status:      "Forbidden",
// 				Description: err.Error(),
// 				Data:        nil,
// 			})
// 			c.Abort()
// 			return
// 		}

// 		parts := strings.Split(user.Can, "|")
// 		exist := false
// 		for _, part := range parts {
// 			if part == permission {
// 				exist = true
// 			}
// 		}

// 		if exist {
// 			c.Next()
// 		} else {
// 			c.JSON(http.Forbidden.Code, http.Response{
// 				Status:      "Forbidden",
// 				Description: "Permission denied",
// 				Data:        nil,
// 			})
// 			c.Abort()
// 			return
// 		}

// 	}
// }
