package Logger

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var loggerCache sync.Map

const logsBaseDir = "Logs"

// getLogLevel 直接读取 Config/web.toml 的 env_mode，避免循环依赖
func getLogLevel() string {
	cwd, err := os.Getwd()
	if err != nil {
		return "info"
	}
	root := cwd
	for {
		if _, err := os.Stat(filepath.Join(root, "go.mod")); err == nil {
			break
		}
		parent := filepath.Dir(root)
		if parent == root {
			return "info"
		}
		root = parent
	}
	tomlPath := filepath.Join(root, "Config", "web.toml")
	data, err := os.ReadFile(tomlPath)
	if err != nil {
		return "info"
	}
	re := regexp.MustCompile(`(?m)^\s*env_mode\s*=\s*"([^"]*)"`)
	matches := re.FindSubmatch(data)
	if len(matches) < 2 {
		return "info"
	}
	mode := strings.ToLower(strings.TrimSpace(string(matches[1])))
	if mode == "debug" {
		return "debug"
	}
	return "info"
}

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

	mu        sync.Mutex
	writers   map[string]zapcore.WriteSyncer
	lastDate  string
	stopClean chan struct{}
}

func newDailyWriteSyncer(module string, maxSize, maxBackups, maxAge int) *dailyWriteSyncer {
	w := &dailyWriteSyncer{
		module:     module,
		maxSize:    maxSize,
		maxBackups: maxBackups,
		maxAge:     maxAge,
		writers:    make(map[string]zapcore.WriteSyncer),
		stopClean:  make(chan struct{}),
	}
	// 启动定时清理协程
	go w.startCleanupRoutine()
	return w
}

func (w *dailyWriteSyncer) currentWriter() (zapcore.WriteSyncer, error) {
	date := nowFunc().Format("20060102")

	w.mu.Lock()
	defer w.mu.Unlock()

	// 检查日期是否变化
	if w.lastDate != "" && w.lastDate != date {
		// 关闭旧的写入器
		if oldWriter, ok := w.writers[w.lastDate]; ok {
			if closer, ok := oldWriter.(io.Closer); ok {
				closer.Close()
			}
			delete(w.writers, w.lastDate)
		}
	}

	// 更新最后日期
	w.lastDate = date

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

// startCleanupRoutine 启动定时清理协程，定期检查并清理过期的写入器
func (w *dailyWriteSyncer) startCleanupRoutine() {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			w.cleanupOldWriters()
		case <-w.stopClean:
			return
		}
	}
}

// cleanupOldWriters 清理过期的写入器（只保留当天的写入器）
func (w *dailyWriteSyncer) cleanupOldWriters() {
	w.mu.Lock()
	defer w.mu.Unlock()

	today := nowFunc().Format("20060102")
	for date, writer := range w.writers {
		if date != today {
			if closer, ok := writer.(io.Closer); ok {
				closer.Close()
			}
			delete(w.writers, date)
		}
	}
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
	// 停止定时清理协程
	close(w.stopClean)

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
		Level:      getLogLevel(),
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
	}

	// 创建独立的 logger 实例（不修改全局变量）
	writeSyncer := newDailyWriteSyncer(name, lc.MaxSize, lc.MaxBackups, lc.MaxAge)
	localLogger, err := createMultiOutputLogger(lc, writeSyncer)
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
