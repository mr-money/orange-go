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
	//用户信息
	apiGroup := r.Group("/api")

	user := apiGroup.Group("/user")
	{
		user.GET("/userInfo", User.GetUserInfo)
		user.GET("/userList", User.GetUserListPage)
		user.GET("/add", User.Add)
	}
}
