package controller

import (
	"net/http"
	"org/otaku/blog/dto"
	"org/otaku/blog/rest"
	"org/otaku/blog/router"
	"org/otaku/blog/service"

	"github.com/gin-gonic/gin"
)

//登录
func login(c *gin.Context) {
	loginReq := dto.LoginReq{}
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		rest.Reply400Error(c, rest.ParamError, err.Error())
		return
	}
	//调用service登录
	accessToken, err := service.LoginService.Login(&loginReq)
	if err != nil {
		rest.Reply400Error(c, rest.BussinessError, err.Error())
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func register(c *gin.Context) {
	registerReq := dto.RegisterReq{}
	if err := c.ShouldBindJSON(&registerReq); err != nil {
		rest.Reply400Error(c, rest.ParamError, err.Error())
		return
	}
	if err := service.LoginService.Register(&registerReq); err != nil {
		rest.Reply400Error(c, rest.BussinessError, err.Error())
	}
}

func init() {
	router.AddRouterConfig("/login", router.RouteConfig{
		HandlerFunc:       login,
		Method:            "POST",
		NeedAuthenticated: false,
	})
	router.AddRouterConfig("/register", router.RouteConfig{
		HandlerFunc:       register,
		Method:            "POST",
		NeedAuthenticated: false,
	})
}
