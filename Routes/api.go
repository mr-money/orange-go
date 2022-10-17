package Routes

import (
	"github.com/gin-gonic/gin"
	"go-study/App/Api/QueueDemo"
	"go-study/App/Api/User"
	"go-study/MiddleWare"
)

//
// Api
// @Description: api路由
// @param r
// @return *gin.Engine
//
func Api(r *gin.Engine) {
	apiGroup := r.Group(
		"/api", //api组
		//MiddleWare.CSRF(),      //验证csrf
		//MiddleWare.CSRFToken(), //生成csrf
	)

	// 注册登录
	apiGroup.POST("/register", User.Register)
	apiGroup.POST("/login", User.Login)

	//用户信息
	user := apiGroup.Group("/user")

	//队列测试
	user.GET("/queueTest", QueueDemo.QueueTest)

	user.Use(MiddleWare.Auth())
	{
		user.GET("/userInfo", User.GetUserInfo)
		user.GET("/userList", User.GetUserListPage)
		user.GET("/add", User.Add)
	}

}
