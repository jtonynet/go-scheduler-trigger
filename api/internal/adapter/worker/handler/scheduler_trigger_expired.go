package handler

import (
	"context"
	"fmt"
	"log"
	"time"

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

func (ste *SchedulerTriggerExpired) Run(ctx context.Context) error {
	attempt := 0

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		ch, err := ste.pubSub.Subscribe(ctx)
		if err != nil {
			log.Printf("subscribe error: %v", err)
			time.Sleep(backoff(attempt))
			attempt++
			continue
		}

		attempt = 0

		if err := ste.consume(ctx, ch); err != nil {
			if ctx.Err() != nil {
				return nil
			}

			log.Printf("consume error: %v", err)

			_ = ste.pubSub.Close()

			time.Sleep(backoff(attempt))
			attempt++
		}
	}
}

func (ste *SchedulerTriggerExpired) consume(
	ctx context.Context,
	ch <-chan string,
) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case key, ok := <-ch:
			if !ok {
				return fmt.Errorf("subscription channel closed")
			}
			if err := ste.service.Process(ctx, key); err != nil {
				log.Printf("error processing key %s: %v", key, err)
			}
		}
	}
}

func backoff(attempt int) time.Duration {
	d := time.Second * time.Duration(1<<attempt)
	if d > 30*time.Second {
		return 30 * time.Second
	}
	return d
}
