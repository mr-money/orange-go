package Queue

import (
	"github.com/pkg/errors"
	"go-study/Queue/Worker/Api/QueueDemo"
)

const (
	PrintNameFunc  = "print_name"
	PrintName2Func = "print_name2"
)

var tasksList = map[string]map[string]interface{}{}

//
// initTasks
// @Description: 配置队列及相关消费方法
//
func initTasks() {
	for _, conf := range *confList() {
		tasksList[conf.DefaultQueue] = make(map[string]interface{})
		switch conf.DefaultQueue {
		//区分不同队列任务
		case "go_study":
			tasksList[conf.DefaultQueue][PrintNameFunc] = QueueDemo.PrintName
		case "go_study2":
			tasksList[conf.DefaultQueue][PrintName2Func] = QueueDemo.PrintName
		}
	}
}

//
// getQueueByTask
// @Description: 根据任务消费方法获取队列名称
// @param taskName 消费方法名
// @return string 队列名称
// @return error
//
func getQueueByTask(taskName string) (string, error) {
	for queue, tasks := range tasksList {
		if _, ok := tasks[taskName]; ok {
			return queue, nil
		}
	}
	return "", errors.New("任务：" + taskName + " 队列名不存在")
}
