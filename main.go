package main

import (
	"go-study/Config"
	"go-study/Routes"
)

//
//  main
//  @Description: 入口
//
func main() {
	//加载配置
	var webConfig Config.Web

	Config.Include(webConfig)

	//初始化配置
	Config.Init()

	// 加载路由配置
	Routes.Include(
		Routes.Web, //默认web路由
		Routes.Api) //TODO api路由，需要token中间件验证

	//// 初始化路由
	r := Routes.Init()

	//
	//// 监听端口，默认在8080
	// Run(":8000")
	_ = r.Run()
}
