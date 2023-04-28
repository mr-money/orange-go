package Api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"go-study/Routes"
	"go-study/Test"
	"net/http"
	"testing"
)

var router *gin.Engine

func init() {
	router = gin.New()
	//api模块路由
	Routes.Api(router)
}

func TestPing(t *testing.T) {
	res := Test.SetRequest(router, "get", "/api/ping")

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "pong", res.Body.String())
}
