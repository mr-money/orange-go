package Api

import (
	"go-study/App"
	"go-study/Routes"
	"net/http"
)

//
// Run
// @Description: 默认服务
//
func Run() {
	// 加载路由
	Routes.Include(
		Routes.Web, //默认web路由
		Routes.Api, //api路由，需要token中间件验证
	)

	port := "8080"

	//启动服务
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: Routes.GinEngine,
	}

	//优雅关闭
	App.Shutdown(srv)

	//启动自检
	App.PingServer(port, srv)
}
