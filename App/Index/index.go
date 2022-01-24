package Index

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
