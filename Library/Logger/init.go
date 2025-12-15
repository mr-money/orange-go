package Logger

import (
	"go.uber.org/zap"
	"log/slog"
	"sync"
)

var loggerCache sync.Map

var (
	AppLogger  = MustModuleLogger("app")
)

// @Description: 根据模块创建日志
// @param name
// @return *zap.SugaredLogger
func MustModuleLogger(name string) *zap.SugaredLogger {
	// 用 sync.Map 做简单缓存，线程安全
	if v, ok := loggerCache.Load(name); ok {
		return v.(*zap.SugaredLogger)
	}

	// 默认规则：logs/{name}.log
	lc := logConfig{
		Level:      "info", // 默认级别，想区分就外面传进来
		FileName:   "logs/" + name + ".log",
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
