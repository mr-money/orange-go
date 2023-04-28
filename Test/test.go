package Test

import (
	"net/http"
	"net/http/httptest"
	"strings"
)

//
// SetRequest
// @Description: 测试设置请求
// @param route 路由
// @param method 请求方法 GET/POST
// @param path 请求地址
// @return *httptest.ResponseRecorder
//
func SetRequest(route http.Handler, method string, path string) *httptest.ResponseRecorder {
	method = strings.ToUpper(method)
	request, _ := http.NewRequest(method, path, nil)
	response := httptest.NewRecorder()
	route.ServeHTTP(response, request)
	return response
}
