package Routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Web
// @Description: 默认路由
// @param r
// @return *gin.Engine
func Web(r *gin.Engine) {
	//绑定路由规则，执行的函数
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
}
