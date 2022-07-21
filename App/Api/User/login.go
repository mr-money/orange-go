package User

import (
	"github.com/gin-gonic/gin"
	"github.com/shockerli/cvt"
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
	res, token, loginErr := User.Login(userCondition)
	if loginErr != nil {
		c.JSON(500, gin.H{"msg": loginErr.Error()})
		return
	}

	c.JSON(200, gin.H{
		"userInfo": res,
		"token":    token,
	})

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

	//注册
	user, token, loginErr := User.Register(userInfo)
	if loginErr != nil {
		c.JSON(500, gin.H{"msg": loginErr.Error()})
		return
	}

	c.JSON(200, gin.H{
		"userInfo": user,
		"token":    token,
	})
}

//
// AddUser999
// @Description:新增5万条测试数据
// @param c
//
func AddUser999(c *gin.Context) {
	userInfo := make(map[string]string)
	userInfo["password"] = "123456"

	for i := 40001; i <= 50000; i++ {
		userInfo["name"] = "test-name" + cvt.String(i)

		_, _, loginErr := User.Register(userInfo)

		if loginErr != nil {
			c.JSON(500, gin.H{"msg": loginErr.Error()})
			return
		}

	}

	c.JSON(200, gin.H{
		"userInfo": userInfo,
	})
}
