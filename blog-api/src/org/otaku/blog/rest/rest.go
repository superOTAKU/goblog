package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//错误信息结构体
type ErrorObject struct {
	Code int
	Msg  string
}

var ParamError = ErrorObject{
	Code: 1,
	Msg:  "param error",
}

var AuthenticationError = ErrorObject{
	Code: 2,
	Msg:  "authentication error",
}

var BussinessError = ErrorObject{
	Code: 3,
	Msg:  "bussiness error",
}

func ReplyError(c *gin.Context, status int, errorObject ErrorObject, msg string) {
	errorObj := ErrorObject{}
	errorObj.Code = errorObject.Code
	errorObj.Msg = errorObject.Msg
	if msg != "" {
		errorObj.Msg = msg
	}
	c.JSON(status, gin.H{
		"code": errorObj.Code,
		"msg":  errorObj.Msg,
	})
}

func Reply400Error(c *gin.Context, errorObject ErrorObject, msg string) {
	ReplyError(c, http.StatusBadRequest, errorObject, msg)
}

func Reply403Error(c *gin.Context, errorObject ErrorObject, msg string) {
	ReplyError(c, http.StatusForbidden, errorObject, msg)
}
