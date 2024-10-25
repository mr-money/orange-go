package Handler

import (
	"fmt"
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

//
//
// RandOrderNo
// @Description: 生成随机订单号
// @return string
//
func RandOrderNo() string {
	// 获取当前时间
	currentTime := time.Now()

	// 格式化为年月日时分秒
	timePart := currentTime.Format("20060102150405")

	// 生成一个随机数作为订单号后缀
	rand.Seed(time.Now().UnixNano())
	randomPart := rand.Intn(1000000)

	// 格式化为固定长度的字符串
	return fmt.Sprintf("%s%06d", timePart, randomPart)
}
