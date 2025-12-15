package Logger

import (
	"go.uber.org/zap"
	"log/slog"
)

var AppLogger *zap.SugaredLogger

func init() {
	lc := logConfig{
		Level:      "debug",
		FileName:   "Logs/test.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
	}
	err := initLogger(lc)
	if err != nil {
		slog.Error(err.Error())
	}

	//全局日志结构
	AppLogger = logger.Sugar()
}
