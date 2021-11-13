package router

import (
	"container/list"
	"errors"
	"org/otaku/blog/auth"

	"github.com/gin-gonic/gin"

	"strings"
)

type RouteConfig struct {
	HandlerFunc       gin.HandlerFunc //处理函数
	Path              string          //路径
	Method            string          //请求方法
	NeedAuthenticated bool            //是否需要鉴权
}

//url -> 路由配置
var routerList *list.List

//添加路由配置，由各个模块自己负责添加
func AddRouterConfig(routerConfig RouteConfig) {
	routerList.PushBack(routerConfig)
}

//在这里注册所有的处理器
func InitRouter(router *gin.Engine) error {
	//遍历RouterMap，给必要的handler加上鉴权功能（装饰器模式）
	for item := routerList.Front(); item != nil; item = item.Next() {
		routerConfig := item.Value.(RouteConfig)
		handlerFunc := routerConfig.HandlerFunc
		if routerConfig.NeedAuthenticated {
			handlerFunc = auth.AddAuthIntercepter(handlerFunc)
		}
		switch method := strings.ToLower(routerConfig.Method); method {
		case "get":
			router.GET(routerConfig.Path, handlerFunc)
		case "post":
			router.POST(routerConfig.Path, handlerFunc)
		case "delete":
			router.DELETE(routerConfig.Path, handlerFunc)
		case "put":
			router.DELETE(routerConfig.Path, handlerFunc)
		default:
			return errors.New("not supported method: " + method)
		}
	}
	return nil
}

func init() {
	routerList = list.New()
}
