package Routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"orange-go/App/Api/Log"
	"orange-go/App/Api/QueueDemo"
	"orange-go/App/Api/User"
	"orange-go/MiddleWare"
)

// Api
// @Description: api路由
// @param r
// @return *gin.Engine
func Api(r *gin.Engine) {
	apiGroup := r.Group(
		"/api", //api组
		//MiddleWare.CSRF(),      //验证csrf
		//MiddleWare.CSRFToken(), //生成csrf
	)

	//api http 测试
	apiGroup.GET("/ping", func(context *gin.Context) {
		ping := context.DefaultQuery("ping", "pong")

		context.String(http.StatusOK, ping)
	})

	// 注册登录
	apiGroup.POST("/register", User.Register)
	apiGroup.POST("/login", User.Login)

	//用户信息
	user := apiGroup.Group("/user")

	//队列测试
	user.GET("/queue-test", QueueDemo.QueueTest)
	user.GET("/queue-test2", QueueDemo.QueueTest2)

	user.Use(MiddleWare.Auth())
	{
		user.GET("/userinfo", User.GetUserInfo)
		user.GET("/user-list", User.GetUserListPage)
		user.GET("/add", User.Add)
	}

	//mongoDB查询log
	apiGroup.GET("/mongo-logs", Log.MongoLogs)

	//zap日志
	apiGroup.GET("/zap-logs", Log.ZapLogs)

}
