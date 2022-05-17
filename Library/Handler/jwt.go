package Handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/shockerli/cvt"
	"go-study/Library/MyTime"
	"go-study/Model"
	"time"
)

type ApiLoginClaims struct {
	UserId uint64 `json:"user_id"`
	jwt.StandardClaims
}

const tokenExpire = MyTime.SecondsPerDay //设置过期时间 单位s

var keySecret = []byte("Go-Gin-Study123") //加盐秘钥

//
// ApiLoginToken
// @Description: 登录生成jwt
// @param user
// @return string
// @return error
//
func ApiLoginToken(user Model.User) (string, error) {
	// 创建api登录声明
	claims := ApiLoginClaims{
		// 自定义字段
		user.ID, //用户id
		jwt.StandardClaims{
			Audience:  user.Name,                       // 受众
			ExpiresAt: time.Now().Unix() + tokenExpire, // 失效时间
			Id:        cvt.String(user.Uuid),           // 编号
			IssuedAt:  time.Now().Unix(),               // 签发时间
			Issuer:    "admin",                         // 签发人
			NotBefore: time.Now().Unix(),               // 生效时间
			Subject:   "login",                         // 主题
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(keySecret)
}

//
// ParseToken
// @Description: 解密jwt
// @param token
// @return *jwt.StandardClaims
// @return error
//
func ParseToken(token string) (*ApiLoginClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &ApiLoginClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return keySecret, nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*ApiLoginClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}
