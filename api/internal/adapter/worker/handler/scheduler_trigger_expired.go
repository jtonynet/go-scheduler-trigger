package handler

import (
	"context"
	"log"

	"github.com/jtonynet/go-scheduler-trigger/api/bootstrap"
	"github.com/jtonynet/go-scheduler-trigger/api/internal/adapter/pubSub"
	"github.com/jtonynet/go-scheduler-trigger/api/internal/core/service"
)

type SchedulerTriggerExpired struct {
	pubSub  pubSub.PubSub
	service *service.SchedulerTriggerExpired
}

func NewSchedulerTriggerExpired(
	worker bootstrap.Worker,
) *SchedulerTriggerExpired {
	return &SchedulerTriggerExpired{
		pubSub:  worker.TriggerPubSub,
		service: worker.SchedulerTriggerExpired,
	}
}

func (h *SchedulerTriggerExpired) Run(ctx context.Context) error {
	ch, err := h.pubSub.Subscribe(ctx)
	if err != nil {
		return err
	}

	for key := range ch {
		if err := h.service.Process(ctx, key); err != nil {
			log.Printf("error processing key %s: %v", key, err)
		}
	}

	return nil
}
