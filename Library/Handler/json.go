package Handler

import (
	"encoding/json"
	"github.com/shockerli/cvt"
	"log"
)

// JsonToStruct
// @Description: json字符串转struct
// @param jsonStr
// @param structRes
// @return *struct{}
func JsonToStruct(jsonStr string, structData interface{}) interface{} {
	err := json.Unmarshal([]byte(jsonStr), &structData)
	if err != nil {
		log.Println("JsonToStruct:", err)
		log.Println("json:", jsonStr)
		log.Printf("struct:%+v", structData)
		return structData
	}

	return structData
}

// ToJson
// @Description: 数据格式转json
// @param data 原始格式数据
// @return string json字符串
func ToJson(data interface{}) string {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		log.Println("ToJson:", err)
		return ""
	}

	return cvt.String(jsonStr)
}
