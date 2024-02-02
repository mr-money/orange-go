package Gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"orange-go/Config"
	"os"
	"time"
)

// Mysql 本地数据库链接
var Mysql *gorm.DB

func init() {
	//默认mysql连接
	Mysql = connectMysql()
}

// connectMysql
// @Description: 默认mysql数据库连接
// @return *gorm.DB
func connectMysql() *gorm.DB {
	/*dsn := "root:root@(127.0.0.1:3306)/go_gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		panic(dbErr)
	}*/

	//未配置mysql
	if len(Config.GetFieldByName(Config.Configs.Web.DB, "Host")) == 0 {
		return nil
	}

	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		Config.GetFieldByName(Config.Configs.Web.DB, "User"),
		Config.GetFieldByName(Config.Configs.Web.DB, "Pwd"),
		Config.GetFieldByName(Config.Configs.Web.DB, "Host"),
		Config.GetFieldByName(Config.Configs.Web.DB, "Port"),
		Config.GetFieldByName(Config.Configs.Web.DB, "DbName"),
		Config.GetFieldByName(Config.Configs.Web.DB, "Charset"),
	)

	//慢sql和错误
	loggerConf := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容）
		logger.Config{
			SlowThreshold:             2 * time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Warn,     // 日志级别
			IgnoreRecordNotFoundError: false,           // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,            // 彩色打印
		},
	)

	db, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: loggerConf,
		//AutoMigrate 会自动创建数据库外键约束，您可以在初始化时禁用此功能
		DisableForeignKeyConstraintWhenMigrating: true,
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
