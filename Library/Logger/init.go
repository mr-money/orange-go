package Logger

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var loggerCache sync.Map

const logsBaseDir = "Logs"

var nowFunc = func() time.Time {
	return time.Now().In(time.Local)
}

// SetNowFuncForTest swaps the time source used by daily log routing.
// It returns the previous function so tests can restore it afterward.
func SetNowFuncForTest(fn func() time.Time) func() time.Time {
	previous := nowFunc
	nowFunc = fn
	return previous
}

type loggerEntry struct {
	logger *zap.SugaredLogger
	closer io.Closer
}

var (
	AppLogger  = MustModuleLogger("app")
	TestLogger = MustModuleLogger("test")
)

type dailyWriteSyncer struct {
	module     string
	maxSize    int
	maxBackups int
	maxAge     int

	mu      sync.Mutex
	writers map[string]zapcore.WriteSyncer
}

func newDailyWriteSyncer(module string, maxSize, maxBackups, maxAge int) *dailyWriteSyncer {
	return &dailyWriteSyncer{
		module:     module,
		maxSize:    maxSize,
		maxBackups: maxBackups,
		maxAge:     maxAge,
		writers:    make(map[string]zapcore.WriteSyncer),
	}
}

func (w *dailyWriteSyncer) currentWriter() (zapcore.WriteSyncer, error) {
	date := nowFunc().Format("20060102")

	w.mu.Lock()
	defer w.mu.Unlock()

	if writer, ok := w.writers[date]; ok {
		return writer, nil
	}

	dirPath := filepath.Join(logsBaseDir, date)
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return nil, err
	}

	filePath := filepath.Join(dirPath, w.module+".log")
	writer := getLogWriter(filePath, w.maxSize, w.maxBackups, w.maxAge)
	w.writers[date] = writer
	return writer, nil
}

func (w *dailyWriteSyncer) Write(p []byte) (n int, err error) {
	writer, err := w.currentWriter()
	if err != nil {
		return 0, err
	}
	return writer.Write(p)
}

func (w *dailyWriteSyncer) Sync() error {
	writer, err := w.currentWriter()
	if err != nil {
		return err
	}
	return writer.Sync()
}

func (w *dailyWriteSyncer) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	var firstErr error
	for date, writer := range w.writers {
		if closer, ok := writer.(interface{ Close() error }); ok {
			if err := closer.Close(); err != nil && firstErr == nil {
				firstErr = err
			}
		}
		delete(w.writers, date)
	}

	return firstErr
}

// MustModuleLogger 根据模块创建日志（线程安全，每个模块独立logger）
// @param name 模块名称
// @return *zap.SugaredLogger
func MustModuleLogger(name string) *zap.SugaredLogger {
	// 用 sync.Map 做缓存，线程安全
	if v, ok := loggerCache.Load(name); ok {
		return v.(*loggerEntry).logger
	}

	lc := logConfig{
		Level:      "info",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
	}

	// 创建独立的 logger 实例（不修改全局变量）
	writeSyncer := newDailyWriteSyncer(name, lc.MaxSize, lc.MaxBackups, lc.MaxAge)
	localLogger, err := createLogger(lc, writeSyncer)
	if err != nil {
		slog.Error("init logger", "module", name, "err", err)
		panic(err)
	}

	sugar := localLogger.Sugar()
	loggerCache.Store(name, &loggerEntry{
		logger: sugar,
		closer: writeSyncer,
	})

	return sugar
}

// CloseModuleLogger closes the module logger if it was created by this package.
func CloseModuleLogger(name string) error {
	v, ok := loggerCache.LoadAndDelete(name)
	if !ok {
		return nil
	}

	entry := v.(*loggerEntry)
	if entry.closer == nil {
		return nil
	}

	return entry.closer.Close()
}
