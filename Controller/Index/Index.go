package Index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.String(http.StatusOK, "index page")
}

func Home1(c *gin.Context) {
	c.String(http.StatusOK, "home1 page")
}

func Home2(c *gin.Context) {
	c.String(http.StatusOK, "home2 page")
}
