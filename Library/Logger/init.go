package Logger

import (
	"log/slog"
	"sync"
	"time"

	"go.uber.org/zap"
)

var loggerCache sync.Map

var (
	AppLogger  = MustModuleLogger("app")
	TestLogger = MustModuleLogger("test")
)

// MustModuleLogger 根据模块创建日志（线程安全，每个模块独立logger）
// @param name 模块名称
// @return *zap.SugaredLogger
func MustModuleLogger(name string) *zap.SugaredLogger {
	// 用 sync.Map 做缓存，线程安全
	if v, ok := loggerCache.Load(name); ok {
		return v.(*zap.SugaredLogger)
	}

	// 默认规则：Logs/日期/{name}.log
	lc := logConfig{
		Level:      "info",
		FileName:   "Logs/" + time.Now().Format("20060102") + "/" + name + ".log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
	}

	// 创建独立的 logger 实例（不修改全局变量）
	localLogger, err := createLogger(lc)
	if err != nil {
		slog.Error("init logger", "module", name, "err", err)
		panic(err)
	}

	sugar := localLogger.Sugar()
	loggerCache.Store(name, sugar)

	return sugar
}
