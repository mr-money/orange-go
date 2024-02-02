package Log

import (
	"orange-go/Model"
	"orange-go/Repository/Log"
)

// GetLog
// @Description: 分页获取log
// @return map[string]interface{}
func GetLog(page int64, pageSize int64) map[string]interface{} {
	var total int64         //总条数
	var logList []Model.Log //数据列表

	total = Log.CountLog()

	logList = Log.GetLogPage(page, pageSize)

	return map[string]interface{}{
		"total": total,
		"list":  logList,
	}
}
