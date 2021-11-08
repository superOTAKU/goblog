package auth

import (
	"github.com/gin-gonic/gin"
	"org/otaku/blog/rest"
)

var AuthHeader = "Authentication"

func AddAuthIntercepter(handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		//do auth
		authenticated := false
		authHeader := c.GetHeader(AuthHeader)
		
		if !authenticated {
			rest.ReplyError()
			return
		}
		//other function
		handlerFunc(c)
	}
}
