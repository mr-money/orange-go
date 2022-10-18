package QueueDemo

import (
	"github.com/RichardKnop/machinery/v1/tasks"
	"time"
)

func PrintName(name string) (string, error) {
	//return name, tasks.NewErrRetryTaskLater("error:", 3*time.Second)

	if false { //error 3秒重试
		return name, tasks.NewErrRetryTaskLater("error:", 3*time.Second)
	}
	return name, nil
}
