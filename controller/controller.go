package controller

import (
	"net/http"
	err "patent/util"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code      string      `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	RequestId string      `json:"requestId"`
}

func SendResponse(c *gin.Context, err *err.Error, data interface{}) {
	if err != nil {
		c.PureJSON(err.HttpStatus, Response{
			Code:      err.Code,
			Message:   err.Message,
			Data:      nil,
			RequestId: c.GetHeader("patent"),
		})
	} else {
		c.PureJSON(http.StatusOK, Response{
			Code:      "OK",
			Message:   "OK",
			Data:      data,
			RequestId: c.GetHeader("patent"),
		})
	}
}
