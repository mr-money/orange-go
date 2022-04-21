package Index

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"go-study/Config"
	"go-study/Model"
	"net/http"
)

//
// Home
// @Description: 控制器主页
// @param c
//
func Home(c *gin.Context) {
	c.String(http.StatusOK, "index page")
}

//
// Middle
// @Description: 中间件
// @param c
//
func Middle(c *gin.Context) {
	req := c.Query("request")
	fmt.Println("request:", req)
	// 页面接收
	c.JSON(200, gin.H{"request": req})
}

//
// Conf
// @Description: 读取配置
// @param c
//
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

//
// Database
// @Description: 数据库连接池
// @param c
//
func Database(c *gin.Context) {
	var userInfo Model.User
	//userInfo = User.FindById(1)

	c.JSON(200, gin.H{
		"config": Config.Configs.Web,
		"user":   userInfo,
	})
}
