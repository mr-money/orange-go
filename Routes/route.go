package Routes

import "github.com/gin-gonic/gin"

// Option 路由配置
type option func(*gin.Engine)

var options = []option{}

// GinEngine 全局gin.Engine
var GinEngine *gin.Engine

// Include
// @Description: 注册app的路由配置
// @param opts
//
func Include(opts ...option) {
	options = append(options, opts...)
	GinEngine = newGin()
}

// @Description: 创建gin Engine
// @return r
//
func newGin() (r *gin.Engine) {
	r = gin.New()
	for _, opt := range options {
		opt(r)
	}

	return
}
