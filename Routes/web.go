package Routes

import (
	"github.com/gin-gonic/gin"
	"go-study/App/Index"
	"go-study/MiddleWare"
	"net/http"
)

//
// Web
// @Description: 默认路由
// @param r
// @return *gin.Engine
//
func Web(r *gin.Engine) {
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

	//中间件
	r.GET("/middleware", MiddleWare.Middle(), Index.Middle)

	//配置
	r.GET("/conf", Index.Conf)

	//数据库连接
	r.POST("/database", Index.Database)
}
