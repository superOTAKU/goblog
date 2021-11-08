package controller

import (
	"net/http"
	"org/otaku/blog/rest"
	"org/otaku/blog/router"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getBlog(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		rest.Reply400Error(c, rest.ParamError, "id illegal!")
	}
	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"content": "My first blog!",
	})
}

func init() {
	router.AddRouterConfig("/blog/:id", router.RouteConfig{
		HandlerFunc:       getBlog,
		Method:            "GET",
		NeedAuthenticated: true,
	})
}
