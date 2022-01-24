package main

import (
	"github.com/gin-gonic/gin"
	"go-study/Router"
)

//
//  main
//  @Description: 入口
//
func main() {
	// 创建路由引擎
	r := gin.Default()

	//默认web路由
	Router.Web(r)

	//api路由，需要token中间件验证
	//Router.Api(r)

	// 监听端口，默认在8080
	// Run(":8000")
	_ = r.Run()
}
