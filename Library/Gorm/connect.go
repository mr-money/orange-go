package Gorm

import (
	"fmt"
	"gorm.io/gorm"
	"log/slog"
	"orange-go/Config"
	"sync"
)

var (
	Mysql *gorm.DB //默认本地数据库链接

	once sync.Once
)

func InitConnect() {
	once.Do(func() {
		localMysql() //默认本地数据库链接
	})
}

// @Description: 默认本地mysql连接
func localMysql() {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		Config.Configs.Web.DB.User,
		Config.Configs.Web.DB.Pwd,
		Config.Configs.Web.DB.Host,
		Config.Configs.Web.DB.Port,
		Config.Configs.Web.DB.DbName,
		Config.Configs.Web.DB.Charset,
	)

	Mysql = connectMysql(dsn)

	slog.Info(fmt.Sprintf("Mysql [%s.%s]: Connect Success!",
		Config.Configs.Web.DB.Host,
		Config.Configs.Web.DB.DbName,
	))
}
