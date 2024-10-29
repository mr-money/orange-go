package Log

import (
	"orange-go/Library/Logger"
	"orange-go/Model"
	"orange-go/Repository/Log"
)

// GetMongoLog
// @Description: 分页获取log
// @return map[string]interface{}
func GetMongoLog(page int64, pageSize int64) map[string]interface{} {
	var total int64         //总条数
	var logList []Model.Log //数据列表

	total = Log.CountLog()

	logList = Log.GetLogPage(page, pageSize)

	return map[string]interface{}{
		"total": total,
		"list":  logList,
	}
}

func ZapLogs() {
	info := struct {
		Name    string
		Age     int
		MapName map[string]string
	}{
		Name: "Zap",
		Age:  80,
		MapName: map[string]string{
			"abc": "aaaa",
			"ccc": "wwww",
		},
	}
	Logger.TestLogger.Info(info)
	Logger.AppLogger.Info(info)
}
