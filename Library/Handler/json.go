package Handler

import (
	"encoding/json"
	"fmt"
	"github.com/shockerli/cvt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// JsonToStruct
// @Description: json字符串转struct
// @param jsonStr
// @param structRes
// @return *struct{}
// JsonToStruct 把 jsonStr 填进 structData，兼容数字/字符串混用
func JsonToStruct(jsonStr string, dst interface{}) {
	if jsonStr == "" {
		return
	}
	var raw interface{}
	if err := json.Unmarshal([]byte(jsonStr), &raw); err != nil {
		log.Println("jsontool: first unmarshal:", err)
		return
	}
	// 把 raw 转成目标类型
	v := reflect.ValueOf(dst)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		log.Println("jsontool: dst must be non-nil pointer")
		return
	}
	fixed := convert(raw, v.Elem().Type())
	// 回填
	reflect.ValueOf(dst).Elem().Set(reflect.ValueOf(fixed))
}

// convert 把弱类型值 v 强转成目标类型 t 的值
func convert(v interface{}, t reflect.Type) interface{} {
	if v == nil {
		return reflect.Zero(t).Interface()
	}
	rv := reflect.ValueOf(v)

	// 如果已经是对应类型，直接返回
	if rv.Type().AssignableTo(t) {
		return v
	}

	// 时间类型单独处理
	if t == reflect.TypeOf(time.Time{}) {
		return parseTime(v)
	}

	switch t.Kind() {
	case reflect.Ptr:
		inner := convert(v, t.Elem())
		ptr := reflect.New(t.Elem())
		ptr.Elem().Set(reflect.ValueOf(inner))
		return ptr.Interface()

	case reflect.Struct:
		return convertStruct(v, t)

	case reflect.Slice, reflect.Array:
		return convertSlice(v, t)

	case reflect.Map:
		return convertMap(v, t)

	case reflect.String:
		return toString(v)

	case reflect.Bool:
		return toBool(v)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if n, ok := toInt64(v); ok {
			return reflect.ValueOf(n).Convert(t).Interface()
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if n, ok := toUint64(v); ok {
			return reflect.ValueOf(n).Convert(t).Interface()
		}

	case reflect.Float32, reflect.Float64:
		if f, ok := toFloat64(v); ok {
			return reflect.ValueOf(f).Convert(t).Interface()
		}
	}

	// 默认原样返回
	return v
}
func convertStruct(v interface{}, t reflect.Type) interface{} {
	src, ok := v.(map[string]interface{})
	if !ok {
		return reflect.Zero(t).Interface()
	}
	out := reflect.New(t).Elem()
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		tag := sf.Tag.Get("json")
		name := strings.Split(tag, ",")[0]
		if name == "" || name == "-" {
			name = sf.Name
		}
		val, ok := src[name]
		if !ok {
			continue
		}
		fixed := convert(val, sf.Type)
		out.Field(i).Set(reflect.ValueOf(fixed))
	}
	return out.Interface()
}

func convertSlice(v interface{}, t reflect.Type) interface{} {
	src, ok := v.([]interface{})
	if !ok {
		return reflect.Zero(t).Interface()
	}
	slice := reflect.MakeSlice(t, len(src), len(src))
	for i, val := range src {
		fixed := convert(val, t.Elem())
		slice.Index(i).Set(reflect.ValueOf(fixed))
	}
	return slice.Interface()
}

func convertMap(v interface{}, t reflect.Type) interface{} {
	src, ok := v.(map[string]interface{})
	if !ok {
		return reflect.Zero(t).Interface()
	}
	out := reflect.MakeMapWithSize(t, len(src))
	for k, val := range src {
		fixed := convert(val, t.Elem())
		out.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(fixed))
	}
	return out.Interface()
}

/* ---------- 基础类型转换工具 ---------- */

func toString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(val)
	case nil:
		return ""
	default:
		return fmt.Sprint(val)
	}
}

func toInt64(v interface{}) (int64, bool) {
	switch val := v.(type) {
	case float64:
		return int64(val), true
	case string:
		if n, err := strconv.ParseInt(val, 10, 64); err == nil {
			return n, true
		}
	case bool:
		if val {
			return 1, true
		}
		return 0, true
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(v).Int(), true
	case uint, uint8, uint16, uint32, uint64:
		return int64(reflect.ValueOf(v).Uint()), true
	}
	return 0, false
}

func toUint64(v interface{}) (uint64, bool) {
	switch val := v.(type) {
	case float64:
		return uint64(val), true
	case string:
		if n, err := strconv.ParseUint(val, 10, 64); err == nil {
			return n, true
		}
	case bool:
		if val {
			return 1, true
		}
		return 0, true
	case int, int8, int16, int32, int64:
		return uint64(reflect.ValueOf(v).Int()), true
	case uint, uint8, uint16, uint32, uint64:
		return reflect.ValueOf(v).Uint(), true
	}
	return 0, false
}

func toFloat64(v interface{}) (float64, bool) {
	switch val := v.(type) {
	case float64:
		return val, true
	case string:
		if f, err := strconv.ParseFloat(val, 64); err == nil {
			return f, true
		}
	case bool:
		if val {
			return 1, true
		}
		return 0, true
	case int, int8, int16, int32, int64:
		return float64(reflect.ValueOf(v).Int()), true
	case uint, uint8, uint16, uint32, uint64:
		return float64(reflect.ValueOf(v).Uint()), true
	}
	return 0, false
}

func toBool(v interface{}) bool {
	switch val := v.(type) {
	case bool:
		return val
	case string:
		b, _ := strconv.ParseBool(val)
		return b
	case float64:
		return val != 0
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(v).Int() != 0
	case uint, uint8, uint16, uint32, uint64:
		return reflect.ValueOf(v).Uint() != 0
	default:
		return false
	}
}

// 时间解析，支持 Unix 秒/毫秒、RFC3339、常用格式
func parseTime(v interface{}) time.Time {
	switch val := v.(type) {
	case float64:
		// Unix 秒或毫秒
		sec := int64(val)
		if sec > 1e10 { // 毫秒
			return time.UnixMilli(sec)
		}
		return time.Unix(sec, 0)
	case string:
		// 尝试 RFC3339
		if t, err := time.Parse(time.RFC3339, val); err == nil {
			return t
		}
		// 尝试常用格式
		for _, layout := range []string{
			"2006-01-02 15:04:05",
			"2006-01-02",
			"2006/01/02 15:04:05",
		} {
			if t, err := time.Parse(layout, val); err == nil {
				return t
			}
		}
	}
	return time.Time{}
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
