package Handler

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

//
// HashAndSalt
// @Description: 加密密码
// @param pwd
// @return string
//
func HashAndSalt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

//
// ComparePasswords
// @Description: 验证密码
// @param hashedPwd 已保存加密过的密码
// @param plainPwd 输入的明文密码
// @return bool
//
func ComparePasswords(hashedPwd string, plainPwd string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	if err != nil {
		log.Println(err, hashedPwd, plainPwd)
		return false
	}
	return true
}
