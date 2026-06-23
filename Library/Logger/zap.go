package Logger

import (
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logConfig struct {
	Level      string `json:"level"`       // Level 最低日志等级，DEBUG<INFO<WARN<ERROR<FATAL 例如：info-->收集info等级以上的日志
	FileName   string `json:"file_name"`   // FileName 日志文件位置
	MaxSize    int    `json:"max_size"`    // MaxSize 进行切割之前，日志文件的最大大小(MB为单位)，默认为100MB
	MaxAge     int    `json:"max_age"`     // MaxAge 是根据文件名中编码的时间戳保留旧日志文件的最大天数。
	MaxBackups int    `json:"max_backups"` // MaxBackups 是要保留的旧日志文件的最大数量。默认是保留所有旧日志文件（尽管 MaxAge 可能仍会导致它们被删除。）
}

// ANSI color codes
const (
	colorReset   = "\033[0m"
	colorCyan    = "\033[36m"
	colorBlue    = "\033[34m"
	colorYellow  = "\033[33m"
	colorRed     = "\033[31m"
	colorMagenta = "\033[35m"
)

// levelToColor maps log levels to ANSI colors
func levelToColor(l zapcore.Level) string {
	switch l {
	case zapcore.DebugLevel:
		return colorCyan
	case zapcore.InfoLevel:
		return colorBlue
	case zapcore.WarnLevel:
		return colorYellow
	case zapcore.ErrorLevel:
		return colorRed
	case zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		return colorMagenta
	default:
		return colorReset
	}
}

// colorCore wraps a Core to add color to entire log lines for console output
type colorCore struct {
	zapcore.Core
	enc zapcore.Encoder
	ws  zapcore.WriteSyncer
}

func (c *colorCore) With(fields []zapcore.Field) zapcore.Core {
	return &colorCore{
		Core: c.Core.With(fields),
		enc:  c.enc.Clone(),
		ws:   c.ws,
	}
}

func (c *colorCore) Check(entry zapcore.Entry, checkedEntry *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(entry.Level) {
		return checkedEntry.AddCore(entry, c)
	}
	return checkedEntry
}

func (c *colorCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	buf, err := c.enc.EncodeEntry(entry, fields)
	if err != nil {
		return err
	}
	defer buf.Free()

	color := levelToColor(entry.Level)
	// Write color + buffer + reset
	_, err = c.ws.Write(append(append([]byte(color), buf.Bytes()...), []byte(colorReset)...))
	return err
}

func (c *colorCore) Sync() error {
	return c.ws.Sync()
}

// localTimeEncoder 使用本地时区编码时间
func localTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.In(time.Local).Format("2006-01-02T15:04:05.000Z07:00"))
}

// localTimeConsoleEncoder 控制台使用的时间格式
func localTimeConsoleEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.In(time.Local).Format("2006-01-02 15:04:05.000"))
}

// getEncoder 获取日志编码格式（使用本地时区，JSON格式用于文件）
func getEncoder() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = localTimeEncoder
	encodeConfig.TimeKey = "time"
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}

// getConsoleEncoder 获取控制台编码格式（人类可读）
func getConsoleEncoder() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = localTimeConsoleEncoder
	encodeConfig.TimeKey = "time"
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encodeConfig)
}

// getLogWriter 获取指定文件的日志写入器
// 刷盘由 dailyWriteSyncer 的定时 Sync 统一管理，避免每次写入都触发 fsync
func getLogWriter(filename string, maxsize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxsize,
		MaxAge:     maxAge,
		MaxBackups: maxBackup,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// createLogger 创建仅输出到文件的 Logger 实例
func createLogger(lCfg logConfig, writeSyncer zapcore.WriteSyncer) (*zap.Logger, error) {
	encoder := getEncoder()

	var l = new(zapcore.Level)
	err := l.UnmarshalText([]byte(lCfg.Level))
	if err != nil {
		return nil, err
	}

	core := zapcore.NewCore(encoder, writeSyncer, l)
	return zap.New(core, zap.AddCaller()), nil
}

// createMultiOutputLogger 创建同时输出到文件和控制台的 Logger 实例
func createMultiOutputLogger(lCfg logConfig, fileSyncer zapcore.WriteSyncer) (*zap.Logger, error) {
	var l = new(zapcore.Level)
	err := l.UnmarshalText([]byte(lCfg.Level))
	if err != nil {
		return nil, err
	}

	jsonEncoder := getEncoder()
	consoleEncoder := getConsoleEncoder()

	consoleSyncer := zapcore.Lock(os.Stdout)
	fileCore := zapcore.NewCore(jsonEncoder, fileSyncer, l)
	consoleCore := &colorCore{
		Core: zapcore.NewCore(consoleEncoder, consoleSyncer, l),
		enc:  consoleEncoder,
		ws:   consoleSyncer,
	}

	core := zapcore.NewTee(fileCore, consoleCore)
	return zap.New(core, zap.AddCaller()), nil
}
