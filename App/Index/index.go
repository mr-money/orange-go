package Index

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/shockerli/cvt"
	"go-study/Config"
	"go-study/Model"
	"go-study/Repository/User"
	"net/http"
)

func Index(c *gin.Context) {
	c.String(http.StatusOK, "index page")
}

func Home1(c *gin.Context) {
	c.String(http.StatusOK, "Home1 page")
}

func Home2(c *gin.Context) {
	c.String(http.StatusOK, "Home2 page")
}

func Middle(c *gin.Context) {
	req := c.Query("request")
	fmt.Println("request:", req)
	// 页面接收
	c.JSON(200, gin.H{"request": req})
}

func Conf(c *gin.Context) {
	var webConfig Config.Web

	_, err := toml.DecodeFile("./Config/web.toml", &webConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("config")
	fmt.Println(webConfig)
	c.JSON(200, gin.H{"config": webConfig})

}

func Database(c *gin.Context) {
	var userInfo Model.User
	userInfo = User.FindById(1)

	c.JSON(200, gin.H{
		"config": Config.Configs.Web,
		"user":   userInfo,
	})
}

//
// GetUserInfo
// @Description:
// @param c
// @param user_id 用户id
//
func GetUserInfo(c *gin.Context) {
	userId := cvt.Uint64(c.Query("user_id"))

	var userInfo Model.User
	userInfo = User.FindById(userId)

	c.JSON(200, gin.H{"user_info": userInfo})

}
