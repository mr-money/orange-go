package Handler

import (
	"golang.org/x/crypto/bcrypt"
	"orange-go/Library/Logger"
)

var handlerLogger = Logger.MustModuleLogger("handler")

// HashAndSalt
// @Description: 加密密码
// @param pwd
// @return string
func HashAndSalt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		handlerLogger.Error("hash password error", "error", err)
	}
	return string(hash)
}

// ComparePasswords
// @Description: 验证密码
// @param hashedPwd 已保存加密过的密码
// @param plainPwd 输入的明文密码
// @return bool
func ComparePasswords(hashedPwd string, plainPwd string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	if err != nil {
		handlerLogger.Error("compare password error", "error", err, "hashedPwd", hashedPwd, "plainPwd", plainPwd)
		return false
	}
	return true
}
