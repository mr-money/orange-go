package MiddleWare

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-study/Library/Handler"
	"net/http"
)

//
// Auth
// @Description: api接口检查jwt登录
// @return gin.HandlerFunc
//
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		auth := context.Request.Header.Get("Authorization")

		//判空
		if len(auth) == 0 {
			context.JSON(http.StatusUnauthorized, gin.H{
				"msg": "signature is empty",
			})
			context.Abort()
			return
		}

		// 校验token
		claim, err := Handler.ParseToken(auth)
		fmt.Println(claim)

		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"msg": err.Error(),
			})
			context.Abort()
			return
		}

		context.Next()
	}
}
