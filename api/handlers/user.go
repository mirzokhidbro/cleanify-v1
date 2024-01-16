package handlers

// func (h *Handler) CreateUser(c *gin.Context) {
// 	var user auth_service.CreateUserRequest

// 	err := c.ShouldBindJSON(&user)
// 	if err != nil {
// 		h.handleResponse(c, http.BadRequest, err.Error())
// 		return
// 	}

// 	resp, err := h.services.UserService().CreateUser(
// 		c.Request.Context(),
// 		&user,
// 	)

// 	if err != nil {
// 		h.handleResponse(c, http.GRPCError, err.Error())
// 		return
// 	}

// 	h.handleResponse(c, http.Created, resp)
// }
