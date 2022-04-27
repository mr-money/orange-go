package User

import (
	"github.com/gin-gonic/gin"
	"go-study/Service/User"
)

//登录
func Login(c *gin.Context) {

}

//
// Register
// @Description: 注册
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
