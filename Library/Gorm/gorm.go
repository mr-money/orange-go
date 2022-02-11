package Gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB database instance
var DBInstance *gorm.DB

func init() {
	DBInstance = connect()
}

//
// Connect
// @Description: 数据库连接
// @return *gorm.DB
//
func connect() *gorm.DB {
	dsn := "root:root@(127.0.0.1:3306)/go_gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		panic(dbErr)
	}

	//todo Config.Configs[0] 切片config获取值

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
