package Handler

//
// JudgeType
// @Description: 判断基本数据类型
// @param items
// @return typeList
//
func JudgeType(value interface{}) (typeName string) {

	switch value.(type) {
	default:
	case string:
		typeName = "string"
	case bool:
		typeName = "bool"
	case float32:
		typeName = "float32"
	case float64:
		typeName = "float64"
	case int:
		typeName = "int"
	case int8:
		typeName = "int8"
	case int16:
		typeName = "int16"
	case int32:
		typeName = "int32"
	case int64:
		typeName = "int64"
	case uint:
		typeName = "uint"
	case uint8:
		typeName = "uint8"
	case uint16:
		typeName = "uint16"
	case uint32:
		typeName = "uint32"
	case uint64:
		typeName = "uint64"
	}

	return typeName
}
