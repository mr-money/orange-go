package Log

import (
	"github.com/gin-gonic/gin"
	"github.com/shockerli/cvt"
	"net/http"
	"orange-go/Service/Log"
)

// @Description: 获取mango日志
// @param c
func MongoLogs(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "20")

	logs := Log.GetMongoLog(cvt.Int64(page),
		cvt.Int64(pageSize))

	c.JSON(
		http.StatusOK,
		gin.H{
			"code": http.StatusOK,
			"data": logs,
			"msg":  "成功",
		})

}

// @Description: zap日志
// @param c
func ZapLogs(c *gin.Context) {
	Log.ZapLogs()

	c.JSON(
		http.StatusOK,
		gin.H{
			"code": http.StatusOK,
			"data": "",
			"msg":  "成功",
		})
}
