package Router

import (
	"github.com/gin-gonic/gin"
	"go-study/Controller/Index"
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

	//关联控制器
	r.GET("/index/index", Index.Index)

	//路由组
	group := r.Group("/group")
	{
		group.GET("/home1", Index.Home1)
		group.GET("/home2", Index.Home2)
	}

}
