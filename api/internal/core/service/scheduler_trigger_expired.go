package service

import (
	"context"

	"github.com/jtonynet/go-scheduler-trigger/api/internal/adapter/email"
	"github.com/jtonynet/go-scheduler-trigger/api/internal/adapter/repository"
)

type SchedulerTriggerExpired struct {
	mail          email.Mail
	shadowKeyRepo repository.SchedulerTrigger
}

func NewSchedulerTriggerExpired(
	mail email.Mail,
	shadowKeyRepo repository.SchedulerTrigger,
) *SchedulerTriggerExpired {
	return &SchedulerTriggerExpired{
		mail:          mail,
		shadowKeyRepo: shadowKeyRepo,
	}
}

func (ste *SchedulerTriggerExpired) Process(ctx context.Context, key string) error {
	scheduleDTO, err := ste.shadowKeyRepo.Retrieve(ctx, key)
	if err != nil {
		return err
	}

	if err := ste.mail.Send(
		scheduleDTO.Email,
		"TESTE",
		scheduleDTO.Message,
	); err != nil {
		return err
	}

	return ste.shadowKeyRepo.Delete(ctx, key)
}
