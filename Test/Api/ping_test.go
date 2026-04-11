package Api

import (
	"github.com/go-playground/assert/v2"
	"net/http"
	"orange-go/Test"
	"testing"
)

// TestPing
// @Description: 测试自定义ping参数响应
// @param t
func TestPing(t *testing.T) {
	res := Test.Get("/api/ping", map[string]string{
		"ping": "hello",
	})

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "hello", res.Body.String())
}

// TestPingDefault
// @Description: 测试默认ping响应
// @param t
func TestPingDefault(t *testing.T) {
	res := Test.Get("/api/ping", map[string]string{})

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "pong", res.Body.String())
}
