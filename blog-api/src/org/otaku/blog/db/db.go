package db

import (
	"org/otaku/blog/conf"

	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func init() {
	var err error
	db, err = gorm.Open(mysql.Open(conf.GetStr("DbUrl")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var sqlDB *sql.DB
	sqlDB, err = db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(conf.GetInt("MaxIdle"))
	sqlDB.SetMaxOpenConns(conf.GetInt("MaxOpen"))
}
