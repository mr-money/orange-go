package Index

import (
	"github.com/gin-gonic/gin"
	"github.com/shockerli/cvt"
	"go-study/Model"
	"go-study/Repository/User"
	userSer "go-study/Service/User"
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
	userInfo = User.FindById(userId)

	c.JSON(200, gin.H{"user_info": userInfo})

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
	userName := cvt.String(c.Query("name"))

	userInfo := make(map[string]interface{})
	userInfo["name"] = userName
	res := userSer.CreateUser(userInfo)

	c.JSON(200, gin.H{"res": res})

}
