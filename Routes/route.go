package Routes

import "github.com/gin-gonic/gin"

type Option func(*gin.Engine)

var options = []Option{}

var GinEngine *gin.Engine

//
// Include
// @Description: 注册app的路由配置
// @param opts
//
func Include(opts ...Option) {
	options = append(options, opts...)
	GinEngine = newGin()
}

func newGin() (r *gin.Engine) {
	r = gin.New()
	for _, opt := range options {
		opt(r)
	}

	return
}
