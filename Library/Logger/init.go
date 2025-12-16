package Logger

import (
	"go.uber.org/zap"
	"log/slog"
	"sync"
	"time"
)

var loggerCache sync.Map

var (
	AppLogger = MustModuleLogger("app")
)

// @Description: 根据模块创建日志
// @param name
// @return *zap.SugaredLogger
func MustModuleLogger(name string) *zap.SugaredLogger {
	// 用 sync.Map 做简单缓存，线程安全
	if v, ok := loggerCache.Load(name); ok {
		return v.(*zap.SugaredLogger)
	}

	// 默认规则：logs/日期/{name}.log
	lc := logConfig{
		Level:      "info", // 默认级别，想区分就外面传进来
		FileName:   "Logs/" + time.Now().Format("20060102") + "/" + name + ".log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
	}
	if err := initLogger(lc); err != nil {
		slog.Error("init logger", "module", name, "err", err)
		panic(err)
	}

	sugar := logger.Sugar()
	loggerCache.Store(name, sugar)

	return sugar
}
