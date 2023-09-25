package Handler

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
	err := json.Unmarshal([]byte(jsonStr), structData)
	if err != nil {
		log.Panicln("JsonToStruct:", err)
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
		log.Panicln("ToJson:", err)
		return ""
	}

	return cvt.String(jsonStr)
}

//
// JsonToMap
// @Description: json字符串转map
// @param jsonStr
// @return map[string]string
// @return error
//
func JsonToMap(jsonStr string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
