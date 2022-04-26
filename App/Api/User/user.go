package User

import (
	"github.com/gin-gonic/gin"
	"github.com/shockerli/cvt"
	"go-study/Model"
	"go-study/Service/User"
)

//
// GetUserInfo
// @Description:根据id获取用户
// @param c
// @param user_id 用户id
//
func GetUserInfo(c *gin.Context) {
	userId := cvt.Uint64(c.Query("user_id"))

	var userInfo Model.User
	userInfo = User.FindUser(userId)

	c.JSON(200, gin.H{"user_info": userInfo})

}

//
// GetUserListPage
// @Description: 分页获取用户列表
// @param c
//
func GetUserListPage(c *gin.Context) {
	page := cvt.Uint64(c.Query("page"))
	pageSize := cvt.Uint64(c.Query("page_size"))
	uuid := c.Query("uuid")
	userName := c.Query("user_name")

	var userList []Model.User

	//搜索条件
	search := map[string]interface{}{
		"uuid":      uuid,
		"user_name": userName,
	}

	//用户列表
	userList = User.SelectUserListPage(search, page, pageSize)

	c.JSON(200, gin.H{"userList": userList})

}

//
//
// Add
// @Description: 创建用户
// @param c
// @param string name 用户名称
// @return int 插入数据的主键
//
func Add(c *gin.Context) {
	userName := c.Query("name")

	userInfo := make(map[string]interface{})
	userInfo["name"] = userName
	//res := User.Register(userInfo)

	c.JSON(200, gin.H{"res": userInfo})

}
