package Queue

import (
	"fmt"
	"github.com/RichardKnop/machinery/v1/config"
	"go-study/Config"
)

var ServerConf *config.Config

func initConf() {
	ServerConf = &config.Config{
		Broker: fmt.Sprintf("redis://%s:%s/%s",
			Config.GetFieldByName(Config.Configs.Web.Redis, "Host"),
			Config.GetFieldByName(Config.Configs.Web.Redis, "Port"),
			"1",
		),
		DefaultQueue: "go_study",
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
		/*AMQP: &config.AMQPConfig{
			Exchange:      "go_study",
			ExchangeType:  "direct",
			BindingKey:    "go_study_task",
			PrefetchCount: 3,
		},*/
	}
}
