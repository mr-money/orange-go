package Routes

import (
	"github.com/gin-gonic/gin"
	"go-study/App/Api/User"
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

	//用户信息
	user := apiGroup.Group("/user")
	{
		user.GET("/userInfo", User.GetUserInfo)
		user.GET("/userList", User.GetUserListPage)
		user.GET("/add", User.Add)
		user.POST("/register", User.Register)
	}
}
