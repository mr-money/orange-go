package Handler

//
// EmptyStrDef
// @Description: 字符串判空设置默认值
// @param str 原始字符串
// @param defaults 默认值
// @return string
//
func EmptyStrDef(str *string, defaults string) *string {
	if *str == "" {
		*str = defaults
	}

	return str
}
