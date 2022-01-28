package Gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB database instance
var DBInstance *gorm.DB

func Connect() {
	dsn := "root:root@(127.0.0.1:3306)/go_gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		panic(dbErr)
	}

	DBInstance = db
}
