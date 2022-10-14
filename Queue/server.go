package Queue

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"log"
	"os"
)

var Server *machinery.Server

func InitQueue() {
	rootPath, _ := os.Getwd()
	cnf, err := config.NewFromYaml(rootPath+"/Config/queue.yml", false)
	if err != nil {
		log.Println("config failed", err)
		return
	}

	Server, err = machinery.NewServer(cnf)
	if err != nil {
		log.Println("start server failed", err)
		return
	}

	worker := Server.NewWorker("queue", 1)
	go func() {
		err = worker.Launch()
		if err != nil {
			log.Println("start worker error", err)
			return
		}
	}()
}
