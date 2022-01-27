package Gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() *gorm.DB {
	dsn := "root:root@(127.0.0.1:3306)/go_gin?charset=utf8mb4&parseTime=True&loc=Local"
	var dbErr error
	db, dbErr = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		fmt.Println(dbErr)
		return db
	}

	return db
}
