package Wechat

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

//
// StepInfo
// @Description: 微信运动用户数据
//
type StepInfo struct {
	StepInfoList []struct {
		Timestamp int `json:"timestamp"`
		Step      int `json:"step"`
	} `json:"stepInfoList"`
}

//
// GetWxRunData
// @Description: 解析微信运动数据
// @param sessionKey
// @param encryptedData
// @param iv
//
func GetWxRunData(sessionKey, encryptedData, iv string) (stepInfo StepInfo) {
	// 对密钥和 IV 进行 Base64 解码
	key, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		fmt.Println("Error decoding session key:", err)
		return
	}

	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		fmt.Println("Error decoding IV:", err)
		return
	}

	// 对密文进行 Base64 解码
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		fmt.Println("Error decoding encrypted data:", err)
		return
	}

	// 使用密钥和 IV 创建一个 AES 解密器
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error creating cipher:", err)
		return
	}

	mode := cipher.NewCBCDecrypter(block, ivBytes)

	// 解密数据
	decryptedData := make([]byte, len(ciphertext))
	mode.CryptBlocks(decryptedData, ciphertext)

	// 去除 PKCS#7 填充
	padding := int(decryptedData[len(decryptedData)-1])
	decryptedData = decryptedData[:len(decryptedData)-padding]

	// 解析 JSON
	if err := json.Unmarshal(decryptedData, &stepInfo); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	return stepInfo
}
