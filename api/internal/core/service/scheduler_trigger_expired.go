package service

import (
	"context"
	"log"

	"github.com/jtonynet/go-scheduler-trigger/api/internal/adapter/email"
	"github.com/jtonynet/go-scheduler-trigger/api/internal/adapter/pubSub"
	"github.com/jtonynet/go-scheduler-trigger/api/internal/adapter/repository"
)

type SchedulerTriggerExpired struct {
	mail          email.Mail
	triggerPubSub pubSub.PubSub
	shadowKeyRepo repository.SchedulerTrigger
}

func NewSchedulerTriggerExpired(
	mail email.Mail,
	triggerPubSub pubSub.PubSub,
	shadowKeyRepo repository.SchedulerTrigger,
) *SchedulerTriggerExpired {
	return &SchedulerTriggerExpired{
		mail,
		triggerPubSub,
		shadowKeyRepo,
	}
}

func (ste *SchedulerTriggerExpired) Execute() {
	ctx := context.Background()

	triggerExpiredChannel, err := ste.triggerPubSub.Subscribe(ctx)
	if err != nil {
		log.Fatal("cannot subscribe Trigger channel: ", err)
	}

	for key := range triggerExpiredChannel {
		scheduleDTO, _ := ste.shadowKeyRepo.Retrieve(ctx, key)
		ste.mail.Send(
			scheduleDTO.Email,
			"TESTE",
			scheduleDTO.Message,
		)

		ste.shadowKeyRepo.Delete(ctx, key)
	}
}
