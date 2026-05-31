package Logger

import (
	"io"
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

// syncWriteSyncer 包装 WriteSyncer，确保每次写入都同步
type syncWriteSyncer struct {
	ws     zapcore.WriteSyncer
	closer io.Closer
}

func (s *syncWriteSyncer) Write(p []byte) (n int, err error) {
	n, err = s.ws.Write(p)
	if err == nil {
		_ = s.ws.Sync()
	}
	return n, err
}

func (s *syncWriteSyncer) Sync() error {
	return s.ws.Sync()
}

func (s *syncWriteSyncer) Close() error {
	if s.closer == nil {
		return nil
	}
	return s.closer.Close()
}

// localTimeEncoder 使用本地时区编码时间
func localTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	// 使用本地时区格式化时间，避免硬编码时区偏移
	enc.AppendString(t.In(time.Local).Format("2006-01-02T15:04:05.000Z07:00"))
}

// getEncoder 获取日志编码格式（使用本地时区）
func getEncoder() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = localTimeEncoder
	encodeConfig.TimeKey = "time"
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}

// getLogWriter 获取指定文件的日志写入器（带自动同步）
func getLogWriter(filename string, maxsize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxsize,
		MaxAge:     maxAge,
		MaxBackups: maxBackup,
		Compress:   false,
	}

	// 包装一层，确保每次写入后都自动 sync
	ws := zapcore.AddSync(lumberJackLogger)
	return &syncWriteSyncer{ws: ws, closer: lumberJackLogger}
}

// createLogger 创建独立的 Logger 实例（不修改全局变量）
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
