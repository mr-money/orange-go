package QueueDemo

import (
	"github.com/gin-gonic/gin"
	"go-study/Service/User"
)

// QueueTest 队列测试
func QueueTest(c *gin.Context) {
	userName := c.Query("name")

	res := User.QueueTest(userName)

	c.JSON(200, gin.H{"res": res})
}

func QueueTest2(c *gin.Context) {
	userName := c.Query("name")

	res := User.QueueTest2(userName)

	c.JSON(200, gin.H{"res": res})
}
