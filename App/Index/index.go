package Index

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"orange-go/Library/Logger"
)

var indexLogger = Logger.MustModuleLogger("index")

// Home
// @Description: 控制器主页
// @param c
func Home(c *gin.Context) {
	c.String(http.StatusOK, "index page")
}

// Middle
// @Description: 中间件
// @param c
func Middle(c *gin.Context) {
	req := c.Query("request")
	indexLogger.Debug("request received", "request", req)
	// 页面接收
	c.JSON(200, gin.H{"request": req})
}
