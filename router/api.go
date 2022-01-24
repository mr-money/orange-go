package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//
// Router
// @Description: 注册路由
// @param r
// @return *gin.Engine
//
func Router(r *gin.Engine) {
	//绑定路由规则，执行的函数
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

}
