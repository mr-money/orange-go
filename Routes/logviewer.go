package Routes

import (
	"net/http"
	LogViewerHandler "orange-go/App/LogViewer/Handler"
	"orange-go/Library/Handler"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// LogViewer
// @Description: 日志查看器路由
// @param r
// @return *gin.Engine
func LogViewer(r *gin.Engine) {
	r.Use(LogViewerHandler.DebugModeMiddleware())

	projectRoot, err := Handler.GetProjectRoot()
	if err != nil {
		panic(err)
	}
	r.LoadHTMLGlob(filepath.Join(projectRoot, "App", "LogViewer", "*.html"))

	r.GET("/", LogViewerHandler.Index)

	logviewerGroup := r.Group("/logviewer")
	{
		logviewerGroup.GET("/files", LogViewerHandler.GetFiles)
		logviewerGroup.GET("/files/:date/:name", LogViewerHandler.GetLogFile)
		logviewerGroup.GET("/stream", LogViewerHandler.StreamLog)
	}

	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Page not found")
	})
}
