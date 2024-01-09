package helper

import (
	"net/http"

	response "bw-erp/internal/data/responses"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(ctx *gin.Context, err error) {
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Message: "Error",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
	}
}

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func SuccessResponse(payload interface{}) response.Response {
	Response := response.Response{
		Code:    http.StatusOK,
		Message: "Ok",
		Data:    payload,
	}

	return Response
}

func InArray(val string, array []string) (exists bool, index int) {
	exists = false
	index = -1

	for i, v := range array {
		if val == v {
			index = i
			exists = true
			return
		}
	}

	return
}
