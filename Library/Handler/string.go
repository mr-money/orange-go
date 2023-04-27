package Handler

import (
	"github.com/shockerli/cvt"
	"math/rand"
	"strings"
	"time"
)

//
// RandString
// @Description: 随机字符串
// @param length 字符串长度
// @return []byte
//
func RandString(length int) string {
	var strByte = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	var strByteLen = len(strByte)

	bytes := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		bytes[i] = strByte[r.Intn(strByteLen)]
	}

	return cvt.String(bytes)
}

//
// RemoveChars
// @Description: 删除字符串内指定字符
// @param str 原字符串
// @param charsToRemove 需要删除的字符
//
func RemoveChars(str *string, charsToRemove ...rune) {
	charStr := string(charsToRemove)

	*str = strings.ReplaceAll(*str, charStr, "")
}
