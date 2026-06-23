package Logger

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"
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

// writerSlot 不可变结构体，通过 atomic.Pointer 实现无锁读取
type writerSlot struct {
	date   string
	writer zapcore.WriteSyncer
}

// dailyWriteSyncer 实现跨天自动切分日志文件
// 使用 atomic.Pointer 实现无锁快路径，仅在跨天时短暂加锁
type dailyWriteSyncer struct {
	module     string
	maxSize    int
	maxBackups int
	maxAge     int

	current   atomic.Pointer[writerSlot]
	mu        sync.Mutex // 仅保护 rotate 操作
	stopClean chan struct{}
}

func newDailyWriteSyncer(module string, maxSize, maxBackups, maxAge int) *dailyWriteSyncer {
	w := &dailyWriteSyncer{
		module:     module,
		maxSize:    maxSize,
		maxBackups: maxBackups,
		maxAge:     maxAge,
		stopClean:  make(chan struct{}),
	}
	go w.startSyncRoutine()
	return w
}

// Write 实现 io.Writer，快路径无锁
func (w *dailyWriteSyncer) Write(p []byte) (n int, err error) {
	slot := w.current.Load()
	today := nowFunc().Format("20060102")

	// 快路径：日期未变，无锁写入
	if slot != nil && slot.date == today {
		return slot.writer.Write(p)
	}

	// 慢路径：跨天切分
	return w.rotate(today, p)
}

// rotate 执行跨天切分，double-check 保证只执行一次
func (w *dailyWriteSyncer) rotate(today string, p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	// double-check：其他协程可能已经完成了切分
	if slot := w.current.Load(); slot != nil && slot.date == today {
		return slot.writer.Write(p)
	}

	// 创建新目录和写入器
	dirPath := filepath.Join(logsBaseDir, today)
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return 0, err
	}
	filePath := filepath.Join(dirPath, w.module+".log")
	newWriter := getLogWriter(filePath, w.maxSize, w.maxBackups, w.maxAge)

	// Swap 原子替换，返回旧 writer
	oldSlot := w.current.Swap(&writerSlot{date: today, writer: newWriter})

	// 关闭旧 writer，刷盘剩余数据
	if oldSlot != nil {
		if closer, ok := oldSlot.writer.(io.Closer); ok {
			closer.Close()
		}
	}

	return newWriter.Write(p)
}

// Sync 刷盘当前 writer
func (w *dailyWriteSyncer) Sync() error {
	if slot := w.current.Load(); slot != nil {
		return slot.writer.Sync()
	}
	return nil
}

// startSyncRoutine 定时刷盘，兼顾性能和数据安全
func (w *dailyWriteSyncer) startSyncRoutine() {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			w.Sync()
		case <-w.stopClean:
			return
		}
	}
}

// Close 关闭当前 writer 并停止后台协程
func (w *dailyWriteSyncer) Close() error {
	close(w.stopClean)

	if slot := w.current.Load(); slot != nil {
		_ = slot.writer.Sync()
		if closer, ok := slot.writer.(io.Closer); ok {
			return closer.Close()
		}
	}
	return nil
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
