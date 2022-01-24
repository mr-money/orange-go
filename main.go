package main

import "go-study/Routes"

//
//  main
//  @Description: 入口
//
func main() {
	// 创建路由引擎
	//r := gin.Default()

	// 加载路由配置
	Routes.Include(
		Routes.Web, //默认web路由
		Routes.Api) //TODO api路由，需要token中间件验证

	// 初始化路由
	r := Routes.Init()

	// 监听端口，默认在8080
	// Run(":8000")
	_ = r.Run()
}
