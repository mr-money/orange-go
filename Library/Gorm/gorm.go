package Gorm

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"orange-go/Config"
	"os"
	"time"
)

func init() {
	//环境模式
	gin.SetMode(Config.Configs.Web.Common.EnvModel)

	//初始化连接
	InitConnect()
}

// @Description: 默认mysql数据库连接
// @return *gorm.DB
func connectMysql(dsn string) *gorm.DB {
	/*dsn := "root:root@(127.0.0.1:3306)/go_gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		panic(dbErr)
	}*/

	if dsn == "" {
		panic("Mysql dsn is null")
	}

	//日志
	logLevel := logger.Warn
	if gin.Mode() == "debug" {
		logLevel = logger.Info
	}

	loggerConf := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容）
		logger.Config{
			SlowThreshold:             5 * time.Second, // 慢 SQL 阈值
			LogLevel:                  logLevel,        // 日志级别
			IgnoreRecordNotFoundError: true,            // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,            // 彩色打印
		},
	)

	db, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: loggerConf,
		//AutoMigrate 会自动创建数据库外键约束，您可以在初始化时禁用此功能
		DisableForeignKeyConstraintWhenMigrating: true,
		//全局禁用默认事务
		SkipDefaultTransaction: true,
		//预编译语句
		PrepareStmt: true,
	})
	if dbErr != nil {
		panic(dbErr)
	}

	////连接池
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, _ := db.DB()

	// 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(50)

	// 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(500)

	// 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(1 * time.Hour)

	return db

}
