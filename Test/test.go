package Test

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"orange-go/Routes"
)

var router *gin.Engine

func init() {
	router = gin.New()
	//api模块路由
	Routes.Api(router)
}

// Get
// @Description: 测试GET请求
// @param path 请求地址
// @return *httptest.ResponseRecorder
func Get(uri string, params map[string]string) *httptest.ResponseRecorder {
	data := url.Values{}
	for key, value := range params {
		data.Set(key, value)
	}
	uri = fmt.Sprintf("%s?%s", uri, data.Encode())

	request, _ := http.NewRequest(http.MethodGet, uri, nil)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	return response
}

func Post(uri string, params, headers map[string]string) *httptest.ResponseRecorder {
	requestBody := &bytes.Buffer{}
	writer := multipart.NewWriter(requestBody)

	for key, value := range params {
		_ = writer.WriteField(key, value)
	}
	_ = writer.Close()

	request, _ := http.NewRequest(http.MethodPost, uri, requestBody)

	request.Header.Set("Content-Type", writer.FormDataContentType())
	for key, value := range headers {
		request.Header.Set(key, value)
	}

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	return response
}
