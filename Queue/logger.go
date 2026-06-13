package Queue

import (
	"fmt"

	"orange-go/Library/Logger"

	machineryLog "github.com/RichardKnop/machinery/v1/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapAdapter struct {
	logger *zap.Logger
	level  zapcore.Level
}

func (a *zapAdapter) Print(v ...interface{})                 { a.log(v...) }
func (a *zapAdapter) Printf(format string, v ...interface{}) { a.Logf(format, v...) }
func (a *zapAdapter) Println(v ...interface{})               { a.log(v...) }
func (a *zapAdapter) Fatal(v ...interface{})                 { a.logger.Fatal(fmt.Sprint(v...)) }
func (a *zapAdapter) Fatalf(format string, v ...interface{}) {
	a.logger.Fatal(fmt.Sprintf(format, v...))
}
func (a *zapAdapter) Fatalln(v ...interface{}) { a.logger.Fatal(fmt.Sprint(v...)) }
func (a *zapAdapter) Panic(v ...interface{})   { a.logger.Panic(fmt.Sprint(v...)) }
func (a *zapAdapter) Panicf(format string, v ...interface{}) {
	a.logger.Panic(fmt.Sprintf(format, v...))
}
func (a *zapAdapter) Panicln(v ...interface{}) { a.logger.Panic(fmt.Sprint(v...)) }

func (a *zapAdapter) log(v ...interface{}) {
	msg := fmt.Sprint(v...)
	switch a.level {
	case zapcore.DebugLevel:
		a.logger.Debug(msg)
	case zapcore.WarnLevel:
		a.logger.Warn(msg)
	case zapcore.ErrorLevel:
		a.logger.Error(msg)
	default:
		a.logger.Info(msg)
	}
}

func (a *zapAdapter) Logf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	switch a.level {
	case zapcore.DebugLevel:
		a.logger.Debug(msg)
	case zapcore.WarnLevel:
		a.logger.Warn(msg)
	case zapcore.ErrorLevel:
		a.logger.Error(msg)
	default:
		a.logger.Info(msg)
	}
}

// initMachineryLogger 用 zap 替换 machinery 内部日志，日志级别跟随 env_mode 配置
func initMachineryLogger() {
	sugar := Logger.MustModuleLogger("machinery")
	level := zapcore.InfoLevel
	if sugar.Desugar().Core().Enabled(zapcore.DebugLevel) {
		level = zapcore.DebugLevel
	}

	machineryLog.Set(&zapAdapter{
		logger: sugar.Desugar(),
		level:  level,
	})
}
