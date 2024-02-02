package Queue

import (
	"fmt"
	"github.com/RichardKnop/machinery/v1/config"
	"orange-go/Config"
)

// initConf
// @Description: 初始化队列配置
// @return *config.Config
func initConf() *config.Config {
	return &config.Config{
		DefaultQueue: "go_study", //默认队列名
		//redis队列
		Broker: fmt.Sprintf("redis://%s:%s/%s",
			Config.GetFieldByName(Config.Configs.Web.Redis, "Host"),
			Config.GetFieldByName(Config.Configs.Web.Redis, "Port"),
			"1",
		),
		Redis: &config.RedisConfig{
			MaxIdle:                3,
			IdleTimeout:            240,
			ReadTimeout:            15,
			WriteTimeout:           15,
			ConnectTimeout:         15,
			NormalTasksPollPeriod:  1000,
			DelayedTasksPollPeriod: 500,
		},

		//rabbitMq队列中间件
		/*Broker: fmt.Sprintf("amqp://%s:%s@%s:%s",
			Config.GetFieldByName(Config.Configs.Web.RabbitMq, "User"),
			Config.GetFieldByName(Config.Configs.Web.RabbitMq, "Pwd"),
			Config.GetFieldByName(Config.Configs.Web.RabbitMq, "Host"),
			Config.GetFieldByName(Config.Configs.Web.RabbitMq, "Port"),
		),
		AMQP: &config.AMQPConfig{
			Exchange:      "go_study",
			ExchangeType:  "direct",
			BindingKey:    "go_study_task",
			PrefetchCount: 3,
		},*/
		ResultBackend: fmt.Sprintf("redis://%s:%s/%s",
			Config.GetFieldByName(Config.Configs.Web.Redis, "Host"),
			Config.GetFieldByName(Config.Configs.Web.Redis, "Port"),
			"1",
		),
		ResultsExpireIn: 3600, //结果过期时间
	}
}

// confList
// @Description: 队列配置list
// @return *[]config.Config
func confList() *[]config.Config {
	return &[]config.Config{
		{
			DefaultQueue: "go_study", //队列名
			Broker: fmt.Sprintf("redis://%s:%s/%s",
				Config.GetFieldByName(Config.Configs.Web.Redis, "Host"),
				Config.GetFieldByName(Config.Configs.Web.Redis, "Port"),
				"1",
			),
			ResultBackend: fmt.Sprintf("redis://%s:%s/%s",
				Config.GetFieldByName(Config.Configs.Web.Redis, "Host"),
				Config.GetFieldByName(Config.Configs.Web.Redis, "Port"),
				"1",
			),
			ResultsExpireIn: 3600, //结果过期时间
			Redis: &config.RedisConfig{
				MaxIdle:                3,
				IdleTimeout:            240,
				ReadTimeout:            15,
				WriteTimeout:           15,
				ConnectTimeout:         15,
				NormalTasksPollPeriod:  1000,
				DelayedTasksPollPeriod: 500,
			},
		},
		{
			DefaultQueue: "go_study2", //队列名
			//Broker: "amqp://guest:guest@localhost:5672",
			Broker: fmt.Sprintf("amqp://%s:%s@%s:%s",
				Config.GetFieldByName(Config.Configs.Web.RabbitMq, "User"),
				Config.GetFieldByName(Config.Configs.Web.RabbitMq, "Pwd"),
				Config.GetFieldByName(Config.Configs.Web.RabbitMq, "Host"),
				Config.GetFieldByName(Config.Configs.Web.RabbitMq, "Port"),
			),
			ResultBackend: fmt.Sprintf("redis://%s:%s/%s",
				Config.GetFieldByName(Config.Configs.Web.Redis, "Host"),
				Config.GetFieldByName(Config.Configs.Web.Redis, "Port"),
				"2",
			),
			ResultsExpireIn: 3600, //结果过期时间
			AMQP: &config.AMQPConfig{
				Exchange:      "go_study2",
				ExchangeType:  "direct",
				BindingKey:    "go_study_task",
				PrefetchCount: 3,
			},
		},
	}
}
