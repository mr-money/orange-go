package Queue

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/tasks"
	"go-study/Library/Handler"
	"log"
)

//var server *machinery.Server
var serverMap map[string]*machinery.Server

func Run() {
	/*rootPath, _ := os.Getwd()
	cnf, err := config.NewFromYaml(rootPath+"/Config/queue.yml", false)
	if err != nil {
		log.Println("config failed", err)
		return
	}*/

	/*server, err = machinery.NewServer(initConf())
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
	}()*/

	//初始化队列&消费任务
	initTasks()

	//循环创建队列server
	var err error
	serverMap = make(map[string]*machinery.Server)
	for _, conf := range *confList() {
		serverMap[conf.DefaultQueue], err = machinery.NewServer(&conf)
		if err != nil {
			log.Println("start server: "+conf.DefaultQueue+" failed", err)
			return
		}

		//todo 协程创建队列配置覆盖重复
		worker := serverMap[conf.DefaultQueue].NewWorker(conf.DefaultQueue, 1)
		go func(inWorker *machinery.Worker) {
			err = inWorker.Launch()

			if err != nil {
				log.Println("start worker error", err)
				return
			}
		}(worker)

		//注册任务
		err = serverMap[conf.DefaultQueue].RegisterTasks(tasksList[conf.DefaultQueue])
		if err != nil {
			log.Panicln("register tasks in queue: "+conf.DefaultQueue+" failed", err)
		}
	}

}

// AddTask 加入队列任务
func AddTask(taskName string, params map[string]interface{}) string {
	// 注册任务
	/*err := server.RegisterTask(taskName, tasksList[taskName])
	if err != nil {
		log.Panicln("register task failed", err)
		return ""
	}*/

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
	res, err := asyncResult.Get(1)
	if err != nil {
		log.Panicln(err)
	}
	log.Printf("queue get res is %v\n", tasks.HumanReadableResults(res))
	//return tasks.HumanReadableResults(res)
	return asyncResult.GetState().State
}
