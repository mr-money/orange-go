package Config

import (
	"fmt"
	"github.com/RichardKnop/machinery/v1/config"
)

/*var cnf = &config.Config{
	Broker:        "amqp://guest:guest@localhost:5672/",
	DefaultQueue:  "machinery_tasks",
	ResultBackend: "amqp://guest:guest@localhost:5672/",
	AMQP: &config.AMQPConfig{
		Exchange:     "machinery_exchange",
		ExchangeType: "direct",
		BindingKey:   "machinery_task",
	},
}*/

var RedisQueue = &config.Config{
	DefaultQueue:    "default_queue",
	ResultsExpireIn: 259200,
	Redis: &config.RedisConfig{
		//MaxIdle:                   3,
		//MaxActive:                 3,
		//IdleTimeout:               240,
		//Wait:                      true,
		//ReadTimeout:               15,
		//WriteTimeout:              15,
		//ConnectTimeout:            15,
		//NormalTasksPollPeriod:  1000,
		//DelayedTasksPollPeriod: 500,
		DelayedTasksKey: "default_queue",
	},
}

//初始化队列配置
func init() {
	RedisQueue.Broker = fmt.Sprintf("redis://%s@%s:%s/%s",
		GetFieldByName(Configs.Web.Redis, "Pwd"),  //密码
		GetFieldByName(Configs.Web.Redis, "Host"), //地址
		GetFieldByName(Configs.Web.Redis, "Port"), //端口
		"2", //redis db
	)

	RedisQueue.ResultBackend = fmt.Sprintf("redis://%s@%s:%s/%s",
		GetFieldByName(Configs.Web.Redis, "Pwd"),  //密码
		GetFieldByName(Configs.Web.Redis, "Host"), //地址
		GetFieldByName(Configs.Web.Redis, "Port"), //端口
		"1", //redis db
	)

}
