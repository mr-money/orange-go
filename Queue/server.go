package Queue

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/tasks"
	"go-study/Config"
	"go-study/Library/Handler"
	"log"
)

var server *machinery.Server

// Run 队列服务
func Run() {
	var err error
	server, err = machinery.NewServer(Config.RedisQueue)
	if err != nil {
		log.Println("start server failed", err)
		return
	}

	worker := server.NewWorker("queue", 1)
	go func() {
		err = worker.Launch()
		if err != nil {
			log.Println("start worker error", err)
			return
		}
	}()
}

// AddTask 加入队列任务
func AddTask(taskName string, params map[string]interface{}) string {
	// 注册任务
	err := server.RegisterTasks(taskList)
	if err != nil {
		log.Panicln("register task failed", err)
		return ""
	}

	//构建参数
	var args []tasks.Arg
	for key, param := range params {
		typeName := Handler.JudgeType(param)
		arg := tasks.Arg{
			Name:  key,
			Type:  typeName,
			Value: param,
		}

		args = append(args, arg)
	}

	//参数签名
	signature := &tasks.Signature{
		Name: taskName,
		Args: args,
	}

	//发送任务
	asyncResult, err := server.SendTask(signature)
	if err != nil {
		log.Panicln(err)
	}

	//获取结果
	/*res, err := asyncResult.Get(1)

	if err != nil {
		log.Panicln(err)
	}
	//log.Printf("queue get res is %v\n", tasks.HumanReadableResults(res))*/

	return asyncResult.GetState().State
}
