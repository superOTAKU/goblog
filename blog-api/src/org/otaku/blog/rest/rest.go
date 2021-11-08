package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//错误信息结构体
type ErrorObject struct {
	Status int
	Code   int
	Msg    string
}

var ParamError = ErrorObject{
	Status: http.StatusBadRequest,
	Code:   1,
	Msg:    "param error",
}

func ReplyError(c *gin.Context, errorObject ErrorObject) {
	c.JSON(errorObject.Status, gin.H{
		"code": errorObject.Code,
		"msg":  errorObject.Msg,
	})
}

func ReplyParamError(c *gin.Context, msg string) {
	var errorObject = ParamError
	errorObject.Msg = msg
	ReplyError(c, errorObject)
}
