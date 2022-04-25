package Func

import (
	"github.com/shockerli/cvt"
	"math/rand"
	"time"
)

var strByte = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
var strByteLen = len(strByte)

//
// RandString
// @Description: 随机字符串
// @param length 字符串长度
// @return []byte
//
func RandString(length int) string {

	bytes := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		bytes[i] = strByte[r.Intn(strByteLen)]
	}

	return cvt.String(bytes)
}
