package QueueDemo

import (
	"fmt"
	"time"
)

// PrintNameDelay 队列消费 打印名称
func PrintNameDelay(name string) (string, error) {
	//error 3秒重试
	time.Sleep(10 * time.Second)
	//return name, tasks.NewErrRetryTaskLater("error:", 3*time.Second)
	fmt.Println("delay:" + name)
	return name, nil
}

func PrintName(name string) (string, error) {
	//error 3秒重试
	//return name, tasks.NewErrRetryTaskLater("error:", 3*time.Second)
	fmt.Println("now:" + name)
	return name, nil
}
