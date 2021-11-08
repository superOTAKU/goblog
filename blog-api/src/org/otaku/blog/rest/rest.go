package rest

import (
	"github.com/gin-gonic/gin"
)

//错误信息结构体
type ErrorObject struct {
	Status int
	Code   int
	Msg    string
}

func ReplyError(c *gin.Context, errorObject ErrorObject) {
	c.JSON(errorObject.Status, gin.H{
		"code": errorObject.Code,
		"msg":  errorObject.Msg,
	})
}
