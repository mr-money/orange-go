package Gorm

import (
	"fmt"
	"go-study/Config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// Mysql 本地数据库链接
var Mysql *gorm.DB

func init() {
	//默认mysql连接
	Mysql = connectMysql()

	//AutoMigrate 数据迁移
	migration()
}

//
// connectMysql
// @Description: 默认mysql数据库连接
// @return *gorm.DB
//
func connectMysql() *gorm.DB {
	/*dsn := "root:root@(127.0.0.1:3306)/go_gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		panic(dbErr)
	}*/

	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s%s?charset=%s&parseTime=true&loc=Local",
		Config.GetFieldByName(Config.Configs.Web, "DB", "User"),
		Config.GetFieldByName(Config.Configs.Web, "DB", "Pwd"),
		Config.GetFieldByName(Config.Configs.Web, "DB", "Host"),
		Config.GetFieldByName(Config.Configs.Web, "DB", "Port"),
		Config.GetFieldByName(Config.Configs.Web, "DB", "Prefix"),
		Config.GetFieldByName(Config.Configs.Web, "DB", "DbName"),
		Config.GetFieldByName(Config.Configs.Web, "DB", "Charset"),
	)

	loggerConf := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             2 * time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info,     // 日志级别
			IgnoreRecordNotFoundError: false,           // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,            // 彩色打印
		},
	)

	db, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: loggerConf,
	})
	if dbErr != nil {
		panic(dbErr)
	}

	////连接池
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, _ := db.DB()

	// SetMaxOpenConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	return db

}

//todo AutoMigrate自动建表 https://blog.csdn.net/qq_39787367/article/details/112567822
func migration() {
	fmt.Println("Mysql Migration begin!")

	dbSetTableOptions("用户表", "InnoDB")
}

//
// dbSetTableOptions
// @Description: 表默认设置
// @param comment 表注释
// @param engine 表引擎
// @return *gorm.DB
//
func dbSetTableOptions(engine string, comment string) *gorm.DB {
	//设置表引擎和表注释
	setValue := fmt.Sprintf("ENGINE=%s COMMENT='%s'", engine, comment)

	fmt.Println(setValue)

	return Mysql.Set("gorm:table_options", setValue)
}
