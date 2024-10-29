package main

import (
	"github.com/gin-gonic/gin"
	"orange-go/App/Api"
	"orange-go/Config"
	"orange-go/Database"
	"orange-go/Queue"
)

// main
// @Description: 入口
func main() {
	//环境模式
	gin.SetMode(Config.Configs.Web.Common.EnvModel)

	//数据库迁移
	Database.InitMigrate()

	//队列服务
	Queue.Run()

	//web Api服务 web服务需最后启动
	Api.Run()
}
