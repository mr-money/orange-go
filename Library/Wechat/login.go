package Wechat

import (
	"crypto/tls"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"net/http"
	"orange-go/Library/Handler"
	"strings"
)

// Code2SessionResponse
// @Description: 小程序登录返回
type Code2SessionResponse struct {
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	Errmsg     string `json:"errmsg"`
	Openid     string `json:"openid"`
	Errcode    int    `json:"errcode"`
}

// Code2Session
// @Description:小程序登录
// @param code前端获取code
// @return interface{}
// @return error
func Code2Session(code string) (Code2SessionResponse, error) {
	//小程序配置
	config := GetWxappConfig()

	client := resty.New()

	url := "https://api.weixin.qq.com/sns/jscode2session"
	if strings.HasPrefix(url, "https://") {
		verify := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client.SetTransport(verify)
	}

	//发送请求
	res, err := client.R().
		SetQueryParams(map[string]string{
			"appid":      config.Appid,
			"secret":     config.Appsecret,
			"js_code":    code,
			"grant_type": "authorization_code",
		}).
		Get(url)

	if err != nil {
		return Code2SessionResponse{}, err
	}

	var response Code2SessionResponse
	Handler.JsonToStruct(res.String(), &response)

	if response.Errcode != 0 {
		return response, errors.New(response.Errmsg)
	}

	return response, nil
}
