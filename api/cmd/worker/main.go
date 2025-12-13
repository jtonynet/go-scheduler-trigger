package main

import (
	"log"

	"github.com/jtonynet/go-scheduler-trigger/api/bootstrap"
	"github.com/jtonynet/go-scheduler-trigger/api/config"
	"github.com/jtonynet/go-scheduler-trigger/api/internal/core/service"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	worker, err := bootstrap.NewWorker(cfg)
	if err != nil {
		log.Fatal("cannot initiate worker: ", err)
	}

	schedulerTriggerExpired := service.NewSchedulerTriggerExpired(
		worker.Email,
		worker.TriggerPubSub,
		worker.ShadowKeyRepo,
	)
	schedulerTriggerExpired.Execute()
}
