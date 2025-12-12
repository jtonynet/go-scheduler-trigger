package main

import (
	"context"
	"log"

	"github.com/jtonynet/go-scheduler-trigger/api/bootstrap"
	"github.com/jtonynet/go-scheduler-trigger/api/config"
)

func main() {
	ctx := context.Background()

	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	worker, err := bootstrap.NewWorker(cfg)
	if err != nil {
		log.Fatal("cannot initiate worker: ", err)
	}

	TriggerExpiredChannel, err := worker.TriggerPubSub.Subscribe(ctx)
	if err != nil {
		log.Fatal("cannot subscribe Trigger channel: ", err)
	}

	for key := range TriggerExpiredChannel {
		scheduleDTO, _ := worker.ShadowKeyRepo.Retrieve(ctx, key)
		worker.Email.Send(
			scheduleDTO.Email,
			"TESTE",
			scheduleDTO.Message,
		)

		worker.ShadowKeyRepo.Delete(ctx, key)
	}

}
