package Queue

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/tasks"
	"log"
	"orange-go/Library/Handler"
)

// server对应不同队列queue
var serverMap map[string]*machinery.Server

func Run() {
	/*rootPath, _ := os.Getwd()
	cnf, err := config.NewFromYaml(rootPath+"/Config/queue.yml", false)
	if err != nil {
		log.Println("config failed", err)
		return
	}*/

	var err error

	//初始化队列配置
	conf := initConf()

	//server需要使用map对应不同队列queue
	serverMap = make(map[string]*machinery.Server)

	//循环队列&创建监听消费
	for _, queue := range *getQueues() {
		serverMap[queue.queueName], err = machinery.NewServer(conf)
		if err != nil {
			log.Println("start server failed", err)
			return
		}

		//匿名闭包使用队列组
		func(queueIn queueGroups) {
			//worker := server.NewWorker(queueIn.queueName, 1)
			//创建custom queue
			worker := serverMap[queueIn.queueName].NewCustomQueueWorker(queueIn.queueName, 1, queueIn.queueName)

			//启动goroutine运行worker
			go func(workerIn *machinery.Worker, queueName string) {
				err = workerIn.Launch()
				if err != nil {
					log.Println("start "+queueName+": worker error", err)
					return
				}

			}(worker, queueIn.queueName)

			//注册任务
			err = serverMap[queueIn.queueName].RegisterTasks(queueIn.tasks)
			if err != nil {
				log.Panicln("register tasks in queue: "+queueIn.queueName+" failed", err)
			}
		}(queue)

	}

}

// AddTask 加入队列任务
func AddTask(taskName string, params map[string]interface{}) string {
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

	//获取队列名
	queueName, err := getQueueByTask(taskName)
	if err != nil {
		log.Panicln(err)
	}

	//参数签名
	signature := &tasks.Signature{
		RoutingKey: queueName,

		Name: taskName,
		Args: args,
	}

	//发送任务
	server := serverMap[queueName]
	asyncResult, err := server.SendTask(signature)
	if err != nil {
		log.Panicln(err)
	}

	//获取结果
	/*res, err := asyncResult.Get(1)
	if err != nil {
		log.Panicln(err)
	}
	log.Printf("queue get res is %v\n", tasks.HumanReadableResults(res))*/
	//return tasks.HumanReadableResults(res)

	return asyncResult.GetState().State
}
