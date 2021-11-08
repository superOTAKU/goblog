package router

import (
	"errors"
	"org/otaku/blog/auth"

	"github.com/gin-gonic/gin"

	"strings"
)

type RouteConfig struct {
	HandlerFunc       gin.HandlerFunc //处理函数
	Method            string          //请求方法
	NeedAuthenticated bool            //是否需要鉴权
}

//url -> 路由配置
var routerMap map[string]RouteConfig

//添加路由配置，由各个模块自己负责添加
func AddRouterConfig(path string, routerConfig RouteConfig) {
	routerMap[path] = routerConfig
}

//在这里注册所有的处理器
func InitRouter(router *gin.Engine) error {
	//遍历RouterMap，给必要的handler加上鉴权功能（装饰器模式）
	for path := range routerMap {
		routerConfig := routerMap[path]
		handlerFunc := routerConfig.HandlerFunc
		if routerConfig.NeedAuthenticated {
			handlerFunc = auth.AddAuth(handlerFunc)
		}
		switch method := strings.ToLower(routerConfig.Method); method {
		case "get":
			router.GET(path, handlerFunc)
		case "post":
			router.POST(path, handlerFunc)
		case "delete":
			router.DELETE(path, handlerFunc)
		case "put":
			router.DELETE(path, handlerFunc)
		default:
			return errors.New("not supported method: " + method)
		}
	}
	return nil
}

func init() {
	routerMap = make(map[string]RouteConfig)
}
