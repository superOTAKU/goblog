package main

import (
	"github.com/gin-gonic/gin"

	"org/otaku/blog/router"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	r.Run()
}
