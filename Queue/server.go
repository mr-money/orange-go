package Queue

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/tasks"
	"go-study/Library/Handler"
	"log"
)

var server *machinery.Server

func Run() {
	/*rootPath, _ := os.Getwd()
	cnf, err := config.NewFromYaml(rootPath+"/Config/queue.yml", false)
	if err != nil {
		log.Println("config failed", err)
		return
	}*/

	//初始化队列&消费任务
	/*initTasks()

	//循环创建队列server
	serverMap = make(map[string]*machinery.Server)
	for _, conf := range *confList() {
		func(confIn config.Config) {
			server, err := machinery.NewServer(&confIn)
			if err != nil {
				log.Println("start server failed", err)
				return
			}

			worker := server.NewWorker(confIn.DefaultQueue, 1)
			go func(workerIn *machinery.Worker) {
				err = workerIn.Launch()
				if err != nil {
					log.Println("start "+confIn.DefaultQueue+": worker error", err)
					return
				}

			}(worker)

			//注册任务
			err = server.RegisterTasks(tasksList[confIn.DefaultQueue])
			if err != nil {
				log.Panicln("register tasks in queue: "+confIn.DefaultQueue+" failed", err)
			}

			serverMap[confIn.DefaultQueue] = server
		}(conf)
	}*/

	var err error
	//创建队列server
	conf := initConf()
	server, err = machinery.NewServer(conf)
	if err != nil {
		log.Println("start server failed", err)
		return
	}

	//循环队列&创建监听消费
	for _, queue := range *getQueues() {
		func(queueIn queueGroups) {
			//worker := server.NewWorker(queueIn.queueName, 1)
			worker := server.NewCustomQueueWorker(queueIn.queueName, 1, queueIn.queueName)

			go func(workerIn *machinery.Worker, queueName string) {
				err = workerIn.Launch()
				if err != nil {
					log.Println("start "+queueName+": worker error", err)
					return
				}

			}(worker, queueIn.queueName)

			//注册任务
			err = server.RegisterTasks(queueIn.tasks)
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
	server := server
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
