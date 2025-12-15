package Handler

import (
	"encoding/json"
	"github.com/shockerli/cvt"
	"orange-go/Library/Logger"
)

// @Description: json字符串转struct
// @param jsonStr
// @param structRes
// @return *struct{}
func JsonToStruct(jsonStr string, structData interface{}) interface{} {
	if jsonStr == "" {
		return structData
	}
	err := json.Unmarshal([]byte(jsonStr), &structData)
	if err != nil {
		Logger.AppLogger.Error(
			"JsonToStruct",
			"err:", err,
			"json:", jsonStr,
		)

		return structData
	}

	return structData
}

// @Description: 数据格式转json
// @param data 原始格式数据
// @return string json字符串
func ToJson(data interface{}) string {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		Logger.AppLogger.Error(
			"ToJson",
			"err:", err,
			"json:", jsonStr,
		)
		return ""
	}

	return cvt.String(jsonStr)
}
