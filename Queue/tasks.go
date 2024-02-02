package Queue

import (
	"github.com/pkg/errors"
	"orange-go/Queue/Worker/Api/QueueDemo"
)

// 定义队列消费方法名称
const (
	PrintNameFunc  = "print_name"
	PrintName2Func = "print_name2"
)

// getQueueByTask
// @Description: 根据任务消费方法获取队列名称
// @param taskName 消费方法名
// @return string 队列名称
// @return error
func getQueueByTask(taskName string) (string, error) {
	for _, queue := range *getQueues() {
		if _, ok := queue.tasks[taskName]; ok {
			return queue.queueName, nil
		}
	}

	return "", errors.New("任务：" + taskName + " 队列不存在")
}

// 队列组
type queueGroups struct {
	queueName string                 //队列名称
	tasks     map[string]interface{} //队列下任务 任务名称：任务消费方法
}

// 获取队列组配置
func getQueues() *[]queueGroups {

	return &[]queueGroups{
		{
			"queue_test",
			map[string]interface{}{
				PrintNameFunc: QueueDemo.PrintName,
			},
		},
		{
			"queue_test2",
			map[string]interface{}{
				PrintName2Func: QueueDemo.PrintName2,
			},
		},
	}

}
