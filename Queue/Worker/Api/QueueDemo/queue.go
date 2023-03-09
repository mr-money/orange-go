package QueueDemo

import (
	"github.com/RichardKnop/machinery/v1/tasks"
	"time"
)

// PrintName 队列消费 打印名称
func PrintName(name string) (string, error) {
	//error 3秒重试
	return name, tasks.NewErrRetryTaskLater("error:", 3*time.Second)
}
