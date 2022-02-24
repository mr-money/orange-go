package Gorm

import (
	"fmt"
	"go-study/Config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Mysql 本地数据库链接
var Mysql *gorm.DB

func init() {
	Mysql = connectMysql()
}

//
// connectMysql
// @Description: 默认mysql数据库连接
// @return *gorm.DB
//
func connectMysql() *gorm.DB {
	dsn := "root:root@(127.0.0.1:3306)/go_gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		panic(dbErr)
	}

	//todo issue init先跑了数据库连接再跑的配置初始化

	fmt.Println("Web---------", Config.Configs.Web)

	/*	var db *gorm.DB
		dbConf := reflect.ValueOf(Config.Configs[0]).FieldByName("DB")
		dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s%s?charset=%s&parseTime=true&loc=Local",
			dbConf.FieldByName("user"),
			dbConf.FieldByName("pwd"),
			dbConf.FieldByName("host"),
			dbConf.FieldByName("port"),
			dbConf.FieldByName("prefix"),
			dbConf.FieldByName("db_name"),
			dbConf.FieldByName("charest"),
		)
		db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})*/
	/*for _, conf := range Config.Configs {
		ds := reflect.ValueOf(conf)
		dbConf := ds.FieldByName("DB")

		if !dbConf.IsValid() {
			dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s%s?charset=%s&parseTime=true&loc=Local",
				dbConf.FieldByName("user"),
				dbConf.FieldByName("pwd"),
				dbConf.FieldByName("host"),
				dbConf.FieldByName("port"),
				dbConf.FieldByName("prefix"),
				dbConf.FieldByName("db_name"),
				dbConf.FieldByName("charest"),
			)

			var dbErr error
			db, dbErr = gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if dbErr != nil {
				panic(dbErr)
			}

		}

	}*/
	return db

}

//todo AutoMigrate自动建表 https://blog.csdn.net/qq_39787367/article/details/112567822
