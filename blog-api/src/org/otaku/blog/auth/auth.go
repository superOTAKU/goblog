package auth

import (
	"org/otaku/blog/rest"
	"org/otaku/blog/util"

	"github.com/gin-gonic/gin"
)

var AuthHeader = "Authentication"

func AddAuthIntercepter(handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		//do auth
		authenticated := false
		authHeader := c.GetHeader(AuthHeader)

		if authHeader != "" {
			claim, err := util.VerifyJWTToken(authHeader)
			if err == nil {
				c.Set("UserId", claim.UserId)
				authenticated = true	
			}
		}

		if !authenticated {
			rest.Reply403Error(c, rest.AuthenticationError, "")
			return
		}
		//other function
		handlerFunc(c)
	}
}
