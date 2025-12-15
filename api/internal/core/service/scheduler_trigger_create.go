package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jtonynet/go-scheduler-trigger/api/internal/adapter/repository"
	"github.com/jtonynet/go-scheduler-trigger/api/internal/core/dto"
)

type SchedulerTriggerCreate struct {
	shadowKeyRepo repository.SchedulerTrigger
	triggerRepo   repository.SchedulerTrigger
}

func NewSchedulerTriggerCreate(
	shadowKeyRepo repository.SchedulerTrigger,
	triggerRepo repository.SchedulerTrigger,
) *SchedulerTriggerCreate {
	return &SchedulerTriggerCreate{
		shadowKeyRepo,
		triggerRepo,
	}
}

func (stc *SchedulerTriggerCreate) Execute(ctx context.Context, scheduleReq dto.SchedulerTriggerReq) (*string, error) {
	scheduleReq.UID = uuid.New()
	key := fmt.Sprintf("schedule:%s", scheduleReq.UID.String())

	// TRIGGER TO SEND MESSAGE
	triggerAt, err := mapUTCDataToTimeDuration(scheduleReq.TriggerAt)
	if err != nil {
		return nil, fmt.Errorf("failed to Map UTC data to TimeDuration: %w", err)
	}

	if err := stc.triggerRepo.Create(
		ctx,
		key,
		nil,
		triggerAt,
	); err != nil {
		return nil, fmt.Errorf("failed to persists Trigger data: %w", err)
	}

	// SHADOW KEY - DATA MESSAGE
	if err = stc.shadowKeyRepo.Create(
		ctx,
		key,
		&scheduleReq,
		nil,
	); err != nil {
		return nil, fmt.Errorf("failed to persists Shadow Key data: %w", err)
	}

	uid := scheduleReq.UID.String()
	return &uid, nil
}
