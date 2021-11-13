package controller

import (
	"net/http"
	"org/otaku/blog/dto"
	"org/otaku/blog/rest"
	"org/otaku/blog/router"
	"org/otaku/blog/service"

	"github.com/gin-gonic/gin"
)

func getBlogList(c *gin.Context) {
	userId := c.GetUint64("UserId")
	blogList, err := service.BlogService.GetBlogList(userId)
	if err != nil {
		rest.Reply400Error(c, rest.BussinessError, "get blog list error")
		return
	}
	c.JSON(http.StatusOK, blogList)
}

func AddBlog(c *gin.Context) {
	var reqDTO dto.BlogReq
	if err := c.ShouldBindJSON(&reqDTO); err != nil {
		rest.Reply400Error(c, rest.ParamError, err.Error())
		return
	}
	userId := c.GetUint64("UserId")
	blogId, err := service.BlogService.AddBlog(userId, &reqDTO)
	if err != nil {
		//TODO 打印错误日志
		rest.Reply400Error(c, rest.BussinessError, "add blog error")
		return
	}
	c.String(http.StatusOK, "%d", blogId)
}

func init() {
	router.AddRouterConfig(router.RouteConfig{
		Path:              "/blog",
		HandlerFunc:       getBlogList,
		Method:            "GET",
		NeedAuthenticated: true,
	})

	router.AddRouterConfig(router.RouteConfig{
		Path:              "/blog",
		HandlerFunc:       AddBlog,
		Method:            "POST",
		NeedAuthenticated: true,
	})
}
