package Config

import "github.com/RichardKnop/machinery/v1/config"

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

var DefaultRedis = &config.Config{
	Broker:          "redis://root:@localhost:6379/1",
	DefaultQueue:    "default_queue",
	ResultBackend:   "mongodb://localhost:27017/go_study",
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
