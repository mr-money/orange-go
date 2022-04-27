package User

import (
	"github.com/gin-gonic/gin"
	"go-study/Service/User"
)

//
// Login
// @Description: 普通用户登录
// @param c
//
func Login(c *gin.Context) {
	userName := c.PostForm("name")
	password := c.PostForm("password")

	if len(userName) < 1 || len(password) < 1 {
		c.JSON(400, gin.H{"msg": "参数错误"})
		return
	}

	userCondition := make(map[string]string)
	userCondition["name"] = userName
	userCondition["password"] = password

	//登录
	res, logErr := User.Login(userCondition)
	if logErr != nil {
		c.JSON(400, gin.H{"msg": logErr.Error()})
		return
	}

	c.JSON(200, gin.H{"res": res})

}

//
// Register
// @Description: 普通用户注册
// @param c
//
func Register(c *gin.Context) {
	userName := c.PostForm("name")
	password := c.PostForm("password")

	if len(userName) < 1 || len(password) < 1 {
		c.JSON(400, gin.H{"msg": "参数错误"})
		return
	}

	userInfo := make(map[string]string)
	userInfo["name"] = userName
	userInfo["password"] = password

	res := User.Register(userInfo)

	c.JSON(200, gin.H{"res": res})
}
