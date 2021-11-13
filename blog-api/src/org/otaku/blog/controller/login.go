package controller

import (
	"net/http"
	"org/otaku/blog/dto"
	"org/otaku/blog/rest"
	"org/otaku/blog/router"
	"org/otaku/blog/service"
	"org/otaku/blog/util"

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

//刷新token，目前简单返回一个新的token
func refreshToken(c *gin.Context) {
	if token, err := util.GenJWTToken(c.GetUint64("UserId")); err == nil {
		c.JSON(http.StatusOK, dto.AccessToken{
			Token: token,
		})
		return
	}
	rest.Reply400Error(c, rest.BussinessError, "")
}

func init() {
	router.AddRouterConfig(router.RouteConfig{
		Path:              "/login",
		HandlerFunc:       login,
		Method:            "POST",
		NeedAuthenticated: false,
	})
	router.AddRouterConfig(router.RouteConfig{
		Path:              "/register",
		HandlerFunc:       register,
		Method:            "POST",
		NeedAuthenticated: false,
	})
	router.AddRouterConfig(router.RouteConfig{
		Path:              "/refreshToken",
		HandlerFunc:       refreshToken,
		Method:            "POST",
		NeedAuthenticated: true,
	})

}
