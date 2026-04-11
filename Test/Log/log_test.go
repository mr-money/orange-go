package Log

import (
	"testing"

	"orange-go/Library/Logger"
	"orange-go/Model"
	"orange-go/Repository/Log"
)

// TestZapLogs
// @Description: 测试zap日志功能
// @param t
func TestZapLogs(t *testing.T) {
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

// TestGetMongoLog
// @Description: 测试获取MongoDB日志
// @param t
func TestGetMongoLog(t *testing.T) {
	page := int64(1)
	pageSize := int64(20)

	var total int64
	var logList []Model.Log

	total = Log.CountLog()
	logList = Log.GetLogPage(page, pageSize)

	if total < 0 {
		t.Errorf("Expected total >= 0, got %d", total)
	}

	if len(logList) > int(pageSize) {
		t.Errorf("Expected logList length <= %d, got %d", pageSize, len(logList))
	}
}
