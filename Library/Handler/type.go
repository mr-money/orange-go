package Handler

type typeMap struct {
	index    int
	typeName string
}

//
// JudgeType
// @Description: 判断基本数据类型
// @param items
// @return typeList
//
func JudgeType(items ...interface{}) (typeList []typeMap) {

	for index, value := range items {
		var typeName string
		switch value.(type) {
		case string:
		default:
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

		typeMap := typeMap{
			index:    index,
			typeName: typeName,
		}

		typeList = append(typeList, typeMap)
	}

	return
}
