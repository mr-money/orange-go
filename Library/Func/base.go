package Func

import "encoding/json"

//
// JsonToStruct
// @Description: json字符串转struct
// @param jsonStr
// @param structRes
// @return *struct{}
//
func JsonToStruct(jsonStr string, structRes interface{}) interface{} {
	_ = json.Unmarshal([]byte(jsonStr), &structRes)

	return structRes
}
