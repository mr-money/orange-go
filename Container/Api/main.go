package Api

import (
	"github.com/gin-gonic/gin"
	"go-study/App/Api"
	"go-study/Config"
	"go-study/Queue"
)

//
//  main
//  @Description: 入口
//
func main() {
	//环境模式
	gin.SetMode(Config.GetFieldByName(Config.Configs.Web, "Common", "EnvModel"))

	//队列服务
	Queue.Run()

	//web Api服务 web服务需最后启动
	Api.Run()
}
