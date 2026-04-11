package main

import (
	"github.com/gin-gonic/gin"
	"orange-go/App/LogViewer"
	"orange-go/Config"
)

// main
// @Description: 日志查看器入口
func main() {
	gin.SetMode(Config.Configs.Web.Common.EnvModel)

	LogViewer.Run()
}
