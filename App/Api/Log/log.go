package Log

import (
	"github.com/gin-gonic/gin"
	"github.com/shockerli/cvt"
	"net/http"
	"orange-go/Service/Log"
)

func Logs(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "20")

	logs := Log.GetLog(cvt.Int64(page),
		cvt.Int64(pageSize))

	c.JSON(
		http.StatusOK,
		gin.H{
			"code": http.StatusOK,
			"data": logs,
			"msg":  "成功",
		})

}
