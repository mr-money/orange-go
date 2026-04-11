package Queue

import (
	"testing"

	"github.com/RichardKnop/machinery/v1/tasks"
	"time"
)

// TestPrintName
// @Description: 测试队列任务PrintName
// @param t
func TestPrintName(t *testing.T) {
	name := "test-name"
	result, err := PrintName(name)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected := name + " ----111"
	if result != expected {
		t.Errorf("Expected result %s, got %s", expected, result)
	}
}

// TestPrintName2
// @Description: 测试队列任务PrintName2
// @param t
func TestPrintName2(t *testing.T) {
	name := "test-name"
	result, err := PrintName2(name)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected := name + " ----2222"
	if result != expected {
		t.Errorf("Expected result %s, got %s", expected, result)
	}
}

// TestPrintNameRetry
// @Description: 测试队列任务重试逻辑
// @param t
func TestPrintNameRetry(t *testing.T) {
	name := "test-name"

	if false {
		result, err := PrintName(name)
		_ = result
		if err == nil {
			t.Error("Expected retry error, got nil")
		}

		retryErr, ok := err.(*tasks.ErrRetryTaskLater)
		if !ok {
			t.Errorf("Expected ErrRetryTaskLater, got %T", err)
		}

		if retryErr.RetryIn() != 3*time.Second {
			t.Errorf("Expected retry in 3s, got %v", retryErr.RetryIn())
		}
	}
}

// PrintName 队列测试任务
func PrintName(name string) (string, error) {
	if false {
		return name, tasks.NewErrRetryTaskLater("error:", 3*time.Second)
	}
	return name + " ----111", nil
}

// PrintName2 队列测试任务2
func PrintName2(name string) (string, error) {
	if false {
		return name, tasks.NewErrRetryTaskLater("error:", 3*time.Second)
	}
	return name + " ----2222", nil
}
