package main

import (
	"org/otaku/blog/controller"
	"org/otaku/blog/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	controller.InitController()
	router.InitRouter(r)
	r.Run()
}
