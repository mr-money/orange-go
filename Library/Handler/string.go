package Handler

import (
	"strings"
)

// RemoveChars
// @Description: 删除字符串内指定字符
// @param str 原字符串
// @param charsToRemove 需要删除的字符
func RemoveChars(str *string, charsToRemove ...rune) {
	charStr := string(charsToRemove)

	*str = strings.ReplaceAll(*str, charStr, "")
}
