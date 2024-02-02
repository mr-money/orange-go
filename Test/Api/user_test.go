package Api

import (
	"github.com/go-playground/assert/v2"
	"net/http"
	"orange-go/Library/Handler"
	"orange-go/Model"
	"orange-go/Test"
	"testing"
)

// TestPing
// @Description: http测试链接
// @param t
func TestPing(t *testing.T) {
	res := Test.Get("/api/ping", map[string]string{
		"ping": "pong",
	})

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "pong", res.Body.String())
}

// TestLogin
// @Description: 登录
// @param t
func TestLogin(t *testing.T) {
	param := map[string]string{
		"name":     "test-name1",
		"password": "123456",
	}

	res := Test.Post("/api/login", param, map[string]string{})

	assert.Equal(t, http.StatusOK, res.Code)

	type loginRes struct {
		Token    string     `json:"token"`
		UserInfo Model.User `json:"userInfo"`
	}

	var response loginRes
	Handler.JsonToStruct(res.Body.String(), &response)

	assert.NotEqual(t, "", response.Token)
	assert.Equal(t, uint64(11), response.UserInfo.ID)
}
