package auth

import (
	"github.com/gin-gonic/gin"
)

func AddAuth(handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		//do auth
		authenticated := false
		if !authenticated {
			return
		}
		//other function
		handlerFunc(c)
	}
}
