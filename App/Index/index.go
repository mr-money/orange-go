package Index

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
