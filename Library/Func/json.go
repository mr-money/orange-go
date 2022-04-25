package Func

import (
	"encoding/json"
	"github.com/shockerli/cvt"
	"log"
)

//
// JsonToStruct
// @Description: json字符串转struct
// @param jsonStr
// @param structRes
// @return *struct{}
//
func JsonToStruct(jsonStr string, structData interface{}) interface{} {
	err := json.Unmarshal([]byte(jsonStr), &structData)
	if err != nil {
		log.Fatal("JsonToStruct:", err)
		return nil
	}

	return structData
}

//
// ToJson
// @Description: 数据格式转json
// @param data 原始格式数据
// @return string json字符串
//
func ToJson(data interface{}) string {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		log.Fatal("ToJson:", err)
		return ""
	}

	return cvt.String(jsonStr)
}
