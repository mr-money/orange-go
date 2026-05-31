package main

import (
	"github.com/gin-gonic/gin"
	"orange-go/App/Api"
	"orange-go/Config"
	"orange-go/Database"
	"orange-go/Library/Logger"
	"orange-go/Queue"
)

var mainLogger = Logger.MustModuleLogger("main")

// main
// @Description: 入口
func main() {
	//环境模式
	gin.SetMode(Config.Configs.Web.Common.EnvModel)

	mainLogger.Info("starting orange-go service...")
	mainLogger.Debug("debug mode enabled")

	//数据库迁移
	Database.InitMigrate()

	//队列服务
	Queue.Run()

	mainLogger.Info("all services initialized, starting web server...")

	//web Api服务 web服务需最后启动
	Api.Run()
}
