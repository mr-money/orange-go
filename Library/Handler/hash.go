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
func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
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
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
