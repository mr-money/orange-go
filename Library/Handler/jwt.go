package Handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/shockerli/cvt"
	"orange-go/Library/MyTime"
	"time"
)

type LoginClaims struct {
	Guid string `json:"user_id"`
	jwt.StandardClaims
}

const tokenExpire = MyTime.SecondsPerDay //设置过期时间 单位s

var keySecret = []byte("Go-Gin-Study123") //加盐秘钥

// LoginToken
// @Description: 登录生成jwt
// @param user
// @return string
// @return error
func LoginToken(guid, name string) (string, error) {
	// 创建api登录声明
	claims := LoginClaims{
		// 自定义字段
		guid,
		jwt.StandardClaims{
			Audience:  name,                                       // 受众
			ExpiresAt: time.Now().Unix() + cvt.Int64(tokenExpire), // 失效时间
			Id:        guid,                                       // 编号
			IssuedAt:  time.Now().Unix(),                          // 签发时间
			Issuer:    "admin",                                    // 签发人
			NotBefore: time.Now().Unix(),                          // 生效时间
			Subject:   "login",                                    // 主题
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(keySecret)
}

// ParseToken
// @Description: 解密jwt
// @param token
// @return *jwt.StandardClaims
// @return error
func ParseToken(token string) (*LoginClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &LoginClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return keySecret, nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*LoginClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}
