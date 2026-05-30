package LogViewer

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

// LogEntry 单条日志记录
type LogEntry struct {
	Time        string                 `json:"time"`
	Level       string                 `json:"level"`
	Message     string                 `json:"message,omitempty"`
	Msg         string                 `json:"msg,omitempty"`
	Caller      string                 `json:"caller,omitempty"`
	ExtraFields map[string]interface{} `json:"extra,omitempty"`
}

// LogFileInfo 日志文件信息
type LogFileInfo struct {
	Date string `json:"date"`
	Name string `json:"name"`
	Size int64  `json:"size"`
}

// LogFilesResponse 日志文件列表响应
type LogFilesResponse struct {
	Dates []string      `json:"dates"`
	Files []LogFileInfo `json:"files"`
}

// ReadLogResponse 读取日志响应
type ReadLogResponse struct {
	Entries         []LogEntry `json:"entries"`
	NextOffset      int        `json:"nextOffset"`
	InitialFileSize int64      `json:"initialFileSize"`
	TotalLines      int        `json:"totalLines"`
}

const logsBaseDir = "Logs"

var (
	streamWatchers = make(map[string][]chan<- LogEntry)
	streamMutex    sync.RWMutex
)

// ListLogDates 获取所有日志日期目录
func ListLogDates() ([]string, error) {
	entries, err := os.ReadDir(logsBaseDir)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, err
	}

	var dates []string
	for _, entry := range entries {
		if entry.IsDir() {
			dates = append(dates, entry.Name())
		}
	}

	sort.Sort(sort.Reverse(sort.StringSlice(dates)))
	return dates, nil
}

// ListLogFiles 获取指定日期的日志文件列表
func ListLogFiles(date string) ([]LogFileInfo, error) {
	dirPath := filepath.Join(logsBaseDir, date)
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			return []LogFileInfo{}, nil
		}
		return nil, err
	}

	var files []LogFileInfo
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".log") {
			info, err := entry.Info()
			if err != nil {
				continue
			}
			files = append(files, LogFileInfo{
				Date: date,
				Name: entry.Name(),
				Size: info.Size(),
			})
		}
	}

	return files, nil
}

// ListAllLogFiles 获取所有日志文件（按日期分组）
func ListAllLogFiles() (*LogFilesResponse, error) {
	dates, err := ListLogDates()
	if err != nil {
		return nil, err
	}

	var allFiles []LogFileInfo
	for _, date := range dates {
		files, err := ListLogFiles(date)
		if err != nil {
			continue
		}
		allFiles = append(allFiles, files...)
	}

	return &LogFilesResponse{
		Dates: dates,
		Files: allFiles,
	}, nil
}

// ReadLogFile 读取日志文件内容（正序分页，基于 initialFileSize 避免文件增长导致重复）
func ReadLogFile(date, name string, level, search string, offset, limit int, initialFileSize int64) (*ReadLogResponse, error) {
	filePath := filepath.Join(logsBaseDir, date, name)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// 首次加载时使用当前文件大小作为快照
	if initialFileSize <= 0 {
		initialFileSize = stat.Size()
	}

	// 文件被截断/轮转时重置
	if stat.Size() < initialFileSize {
		initialFileSize = stat.Size()
	}

	// 只读取 initialFileSize 范围内的内容，忽略后续追加的新数据
	content := make([]byte, initialFileSize)
	if initialFileSize > 0 {
		if _, err := io.ReadFull(file, content); err != nil {
			return nil, err
		}
	}

	allLines := strings.Split(string(content), "\n")

	// 去掉末尾空行（文件末尾换行符产生）
	if len(allLines) > 0 && allLines[len(allLines)-1] == "" {
		allLines = allLines[:len(allLines)-1]
	}

	// 正序遍历：从最旧行到最新行
	entries := make([]LogEntry, 0, limit)
	for i := 0; i < len(allLines); i++ {
		line := allLines[i]
		if strings.TrimSpace(line) == "" {
			continue
		}

		entry, err := parseLogLine(line)
		if err != nil {
			continue
		}

		if level != "" && level != "all" && !strings.EqualFold(entry.Level, level) {
			continue
		}

		if search != "" {
			searchLower := strings.ToLower(search)
			found := strings.Contains(strings.ToLower(entry.Message), searchLower)

			if !found {
				for _, v := range entry.ExtraFields {
					var valueStr string
					switch val := v.(type) {
					case string:
						valueStr = val
					default:
						jsonBytes, _ := json.Marshal(val)
						valueStr = string(jsonBytes)
					}
					if strings.Contains(strings.ToLower(valueStr), searchLower) {
						found = true
						break
					}
				}
			}

			if !found {
				continue
			}
		}

		entries = append(entries, entry)

		if limit > 0 && len(entries) >= limit {
			break
		}
	}

	if entries == nil {
		entries = make([]LogEntry, 0)
	}

	// nextOffset: -1 表示没有更多数据
	nextOffset := -1
	if limit > 0 && len(entries) >= limit {
		nextOffset = offset + limit
	}

	// 统计 initialFileSize 范围内的总行数（非空行）
	totalLines := 0
	for _, line := range allLines {
		if strings.TrimSpace(line) != "" {
			totalLines++
		}
	}

	return &ReadLogResponse{
		Entries:         entries,
		NextOffset:      nextOffset,
		InitialFileSize: initialFileSize,
		TotalLines:      totalLines,
	}, nil
}

// parseLogLine 解析单条日志行
func parseLogLine(line string) (LogEntry, error) {
	var raw map[string]interface{}
	err := json.Unmarshal([]byte(line), &raw)
	if err != nil {
		return LogEntry{}, err
	}

	entry := LogEntry{
		ExtraFields: make(map[string]interface{}),
	}

	// 解析标准字段
	for key, value := range raw {
		switch key {
		case "time":
			if time, ok := value.(string); ok {
				entry.Time = time
			}
		case "level":
			if level, ok := value.(string); ok {
				entry.Level = level
			}
		case "msg":
			if msg, ok := value.(string); ok {
				entry.Msg = msg
				if entry.Message == "" {
					entry.Message = msg
				}
			}
		case "message":
			if message, ok := value.(string); ok {
				entry.Message = message
				if entry.Msg == "" {
					entry.Msg = message
				}
			}
		case "caller":
			if caller, ok := value.(string); ok {
				entry.Caller = caller
			}
		default:
			// 保存额外字段
			entry.ExtraFields[key] = value
		}
	}

	return entry, nil
}

// StreamLogFile 实时监听日志文件
func StreamLogFile(date, name string) (<-chan LogEntry, error) {
	filePath := filepath.Join(logsBaseDir, date, name)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	_, err = file.Seek(0, io.SeekEnd)
	if err != nil {
		file.Close()
		return nil, err
	}

	ch := make(chan LogEntry, 100)

	go func() {
		defer func() {
			file.Close()
			close(ch)
			removeWatcher(filePath, ch)
		}()

		reader := bufio.NewReader(file)

		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					time.Sleep(500 * time.Millisecond)
					continue
				}
				return
			}

			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			entry, err := parseLogLine(line)
			if err != nil {
				continue
			}

			select {
			case ch <- entry:
			default:
			}
		}
	}()

	addWatcher(filePath, ch)
	return ch, nil
}

func addWatcher(filePath string, ch chan<- LogEntry) {
	streamMutex.Lock()
	defer streamMutex.Unlock()
	streamWatchers[filePath] = append(streamWatchers[filePath], ch)
}

func removeWatcher(filePath string, ch chan<- LogEntry) {
	streamMutex.Lock()
	defer streamMutex.Unlock()

	watchers, ok := streamWatchers[filePath]
	if !ok {
		return
	}

	var newWatchers []chan<- LogEntry
	for _, w := range watchers {
		if w != ch {
			newWatchers = append(newWatchers, w)
		}
	}

	if len(newWatchers) == 0 {
		delete(streamWatchers, filePath)
	} else {
		streamWatchers[filePath] = newWatchers
	}
}

// EnsureLogDir 确保日志目录存在
func EnsureLogDir(date string) error {
	dirPath := filepath.Join(logsBaseDir, date)
	return os.MkdirAll(dirPath, 0755)
}

// GetLatestLogFile 获取最新的日志文件
func GetLatestLogFile() (string, string, error) {
	dates, err := ListLogDates()
	if err != nil {
		return "", "", err
	}

	if len(dates) == 0 {
		return "", "", fmt.Errorf("no log directories found")
	}

	latestDate := dates[0]
	files, err := ListLogFiles(latestDate)
	if err != nil {
		return "", "", err
	}

	if len(files) == 0 {
		return "", "", fmt.Errorf("no log files found")
	}

	return latestDate, files[0].Name, nil
}
