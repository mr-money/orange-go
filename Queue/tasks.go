package Queue

import "go-study/Queue/Worker/Api/QueueDemo"

// PrintName 任务名称
const (
	PrintName      = "print_name"
	PrintNameDelay = "print_name_delay"
)

// TaskList 对应任务消费方法
var taskList = map[string]interface{}{
	PrintNameDelay: QueueDemo.PrintNameDelay,
	PrintName:      QueueDemo.PrintName,
}
