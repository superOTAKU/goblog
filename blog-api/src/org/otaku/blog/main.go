package main

import (
	"org/otaku/blog/controller"
	"org/otaku/blog/db"
	"org/otaku/blog/domain"
	"org/otaku/blog/router"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	//测试gorm
	db := db.GetDB()
	db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&domain.Blog{UserId: 1, Content: "First Blog"}).Error; err != nil {
			return err
		}
		return nil
	})
	r := gin.Default()
	controller.InitController()
	router.InitRouter(r)
	r.Run()
}
