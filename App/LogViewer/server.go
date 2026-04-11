package LogViewer

import (
	"net/http"
	"orange-go/App"
	"orange-go/Routes"
)

// Run 启动日志查看器服务
func Run() {
	// 加载路由
	Routes.Include(
		Routes.LogViewer,
	)

	port := "8081"

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: Routes.GinEngine,
	}

	// 优雅关闭
	App.Shutdown(srv)

	// 启动自检
	App.PingServer(port, srv)
}
