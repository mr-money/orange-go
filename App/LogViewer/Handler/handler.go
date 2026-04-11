package Handler

import (
	"net/http"
	"orange-go/Config"
	LogViewerService "orange-go/Service/LogViewer"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DebugModeMiddleware 检查是否为 debug 模式
func DebugModeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if Config.Configs.Web.Common.EnvModel != "debug" {
			c.String(http.StatusForbidden, "Log viewer only available in debug mode")
			c.Abort()
			return
		}
		c.Next()
	}
}

// Index 首页
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// GetFiles 获取日志文件列表
func GetFiles(c *gin.Context) {
	files, err := LogViewerService.ListAllLogFiles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, files)
}

// GetLogFile 读取日志文件内容
func GetLogFile(c *gin.Context) {
	date := c.Param("date")
	name := c.Param("name")

	level := c.DefaultQuery("level", "all")
	search := c.DefaultQuery("search", "")
	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "100")

	offset, _ := strconv.Atoi(offsetStr)
	limit, _ := strconv.Atoi(limitStr)

	logs, err := LogViewerService.ReadLogFile(date, name, level, search, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, logs)
}

// StreamLog 实时日志流 (SSE)
func StreamLog(c *gin.Context) {
	date := c.Query("date")
	name := c.Query("name")

	if date == "" || name == "" {
		var err error
		date, name, err = LogViewerService.GetLatestLogFile()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "No log file specified and no latest log found",
			})
			return
		}
	}

	ch, err := LogViewerService.StreamLogFile(date, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	clientGone := c.Writer.CloseNotify()

	for {
		select {
		case entry, ok := <-ch:
			if !ok {
				return
			}
			c.SSEvent("log", entry)
			c.Writer.Flush()
		case <-clientGone:
			return
		}
	}
}
