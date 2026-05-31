package Log

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"orange-go/Library/Logger"
)

func TestDailyWriteSyncerSwitchesDirectoryByDate(t *testing.T) {
	originalNowFunc := Logger.SetNowFuncForTest(func() time.Time {
		return time.Date(2026, 4, 13, 10, 0, 0, 0, time.Local)
	})
	t.Cleanup(func() {
		Logger.SetNowFuncForTest(originalNowFunc)
	})

	originalWD, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	tempDir := t.TempDir()
	if err := os.Chdir(tempDir); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_ = os.Chdir(originalWD)
	})

	logger := Logger.MustModuleLogger("daily-test")
	t.Cleanup(func() {
		_ = Logger.CloseModuleLogger("daily-test")
	})

	logger.Info("day1")

	Logger.SetNowFuncForTest(func() time.Time {
		return time.Date(2026, 4, 14, 10, 0, 0, 0, time.Local)
	})

	logger.Info("day2")

	day1Path := filepath.Join(tempDir, "Logs", "20260413", "daily-test.log")
	day2Path := filepath.Join(tempDir, "Logs", "20260414", "daily-test.log")

	day1Bytes, err := os.ReadFile(day1Path)
	if err != nil {
		t.Fatalf("read day 1 log: %v", err)
	}
	day2Bytes, err := os.ReadFile(day2Path)
	if err != nil {
		t.Fatalf("read day 2 log: %v", err)
	}

	if !strings.Contains(string(day1Bytes), `"msg":"day1"`) {
		t.Fatalf("day 1 log missing entry: %s", string(day1Bytes))
	}
	if strings.Contains(string(day1Bytes), `"msg":"day2"`) {
		t.Fatalf("day 1 log unexpectedly contains day 2 entry: %s", string(day1Bytes))
	}
	if !strings.Contains(string(day2Bytes), `"msg":"day2"`) {
		t.Fatalf("day 2 log missing entry: %s", string(day2Bytes))
	}
	if strings.Contains(string(day2Bytes), `"msg":"day1"`) {
		t.Fatalf("day 2 log unexpectedly contains day 1 entry: %s", string(day2Bytes))
	}
}
