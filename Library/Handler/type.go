package Handler

import (
	"github.com/pkg/errors"
	"reflect"
)

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

//
// IsEmpty
// @Description: 判断数据是否为空
// @param data
// @return error
//
func IsEmpty(data interface{}) error {
	if data == nil {
		return errors.New("Data is nil")
	}

	value := reflect.ValueOf(data)
	switch value.Kind() {
	case reflect.String:
		if value.String() == "" {
			return errors.New("String is empty")
		}
	case reflect.Slice, reflect.Array, reflect.Map:
		if value.Len() == 0 {
			return errors.New("Slice, Array, or Map is empty")
		}
	case reflect.Ptr:
		if value.IsNil() {
			return errors.New("Pointer is nil")
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if value.Int() == 0 {
			return errors.New("Int is empty")
		}
	case reflect.Float32, reflect.Float64:
		if value.Float() == 0.0 {
			return errors.New("Float is empty")
		}
	case reflect.Bool:
		if !value.Bool() {
			return errors.New("Bool is false")
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if value.Uint() == 0 {
			return errors.New("Uint is empty")
		}
	case reflect.Struct:
		// Check for an empty struct
		zeroValue := reflect.Zero(value.Type())
		if value.Interface() == zeroValue.Interface() {
			return errors.New("Struct is empty")
		}
	}

	return nil
}
