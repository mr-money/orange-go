package Routes

import "github.com/gin-gonic/gin"

type Option func(*gin.Engine)

var options = []Option{}

//
// Include
// @Description: 注册app的路由配置
// @param opts
//
func Include(opts ...Option) {
	options = append(options, opts...)
}

//
// Init
// @Description: 初始化路由
// @return *gin.Engine
//
func Init() *gin.Engine {
	r := gin.New()
	for _, opt := range options {
		opt(r)
	}
	return r
}
