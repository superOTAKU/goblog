package auth

import (
	"github.com/gin-gonic/gin"
)

var AuthHeader = "Authentication"

func AddAuth(handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		//do auth
		authenticated := false
		authHeader := c.GetHeader(AuthHeader)
		if authHeader != "" {
			
		}
		if !authenticated {
			return
		}
		//other function
		handlerFunc(c)
	}
}
