package main

import (
	"context"
	"log"

	"github.com/jtonynet/go-scheduler-trigger/api/bootstrap"
	"github.com/jtonynet/go-scheduler-trigger/api/config"
	"github.com/jtonynet/go-scheduler-trigger/api/internal/adapter/worker/handler"
)

func main() {
	ctx := context.Background()

	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	appWorker, err := bootstrap.NewWorker(cfg)
	if err != nil {
		log.Fatal("cannot initiate worker: ", err)
	}

	worker := handler.NewSchedulerTriggerExpired(*appWorker)
	if err := worker.Run(ctx); err != nil {
		log.Fatal("worker stopped with error: ", err)
	}
}
