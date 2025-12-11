package Logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logConfig struct {
	Level      string `json:"level"`       // Level 最低日志等级，DEBUG<INFO<WARN<ERROR<FATAL 例如：info-->收集info等级以上的日志
	FileName   string `json:"file_name"`   // FileName 日志文件位置
	MaxSize    int    `json:"max_size"`    // MaxSize 进行切割之前，日志文件的最大大小(MB为单位)，默认为100MB
	MaxAge     int    `json:"max_age"`     // MaxAge 是根据文件名中编码的时间戳保留旧日志文件的最大天数。
	MaxBackups int    `json:"max_backups"` // MaxBackups 是要保留的旧日志文件的最大数量。默认是保留所有旧的日志文件（尽管 MaxAge 可能仍会导致它们被删除。）
}

var logger *zap.Logger

// @Description: 负责设置 encoding 的日志格式
// @return zapcore.Encoder
func getEncoder() zapcore.Encoder {
	// 获取一个指定的的EncoderConfig，进行自定义
	encodeConfig := zap.NewProductionEncoderConfig()
	// 打印格式: {"level":"info","ts":1662032576.6267354,"msg":"test","line":1}

	// 设置每个日志条目使用的键。如果有任何键为空，则省略该条目的部分。

	// 序列化时间
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 打印格式：{"level":"info","ts":"2022-09-01T19:43:07.178+0800","msg":"test","line":1}

	encodeConfig.TimeKey = "time"
	// 打印格式：{"level":"info","time":"2022-09-01T19:43:20.558+0800","msg":"test","line":1}

	// 将Level序列化为全大写字符串
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 打印格式：{"level":"INFO","time":"2022-09-01T19:43:41.192+0800","msg":"test","line":1}

	// 以 package/file:行 的格式 序列化调用程序，从完整路径中删除除最后一个目录外的所有目录。
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}

// @Description: 负责日志写入的位置
// @param filename 日志文件名称
// @param maxsize 最大文件尺寸
// @param maxBackup 要保留的旧日志文件的最大数量
// @param maxAge 是根据文件名中编码的时间戳保留旧日志文件的最大天数
// @return zapcore.WriteSyncer
func getLogWriter(filename string, maxsize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,  // 文件位置
		MaxSize:    maxsize,   // 进行切割之前,日志文件的最大大小(MB为单位)
		MaxAge:     maxAge,    // 保留旧文件的最大天数
		MaxBackups: maxBackup, // 保留旧文件的最大个数
		Compress:   false,     // 是否压缩/归档旧文件
	}

	// AddSync 将 io.Writer 转换为 WriteSyncer。
	return zapcore.AddSync(lumberJackLogger)
}

// initLogger 初始化Logger
func initLogger(lCfg logConfig) (err error) {
	// 获取日志写入位置
	writeSyncer := getLogWriter(
		lCfg.FileName,
		lCfg.MaxSize,
		lCfg.MaxBackups,
		lCfg.MaxAge,
	)

	// 获取日志编码格式
	encoder := getEncoder()

	// 获取日志最低等级，即>=该等级，才会被写入。
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(lCfg.Level))
	if err != nil {
		return
	}

	// 创建一个将日志写入 WriteSyncer 的核心。
	core := zapcore.NewCore(encoder, writeSyncer, l)
	logger = zap.New(core, zap.AddCaller())

	// 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	zap.ReplaceGlobals(logger)
	return
}
