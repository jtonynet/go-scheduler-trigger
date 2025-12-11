package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jtonynet/go-scheduler-trigger/api/internal/adapter/database"
	"github.com/jtonynet/go-scheduler-trigger/api/internal/core/dto"
)

type SchedulerTriggerCreate struct {
	// TODO: using `database adapter` instead of a `repository adapter` to simplify the example.
	cacheInMemoDB   database.InMemory
	triggerInMemoDB database.InMemory

	// shadowKeyRepo repository.SchedulerTriggerRedis
	// triggerRepo   repository.SchedulerTriggerRedis
}

func NewSchedulerTriggerCreate(
	cacheInMemoDB database.InMemory,
	triggerInMemoDB database.InMemory,
) *SchedulerTriggerCreate {
	return &SchedulerTriggerCreate{
		cacheInMemoDB,
		triggerInMemoDB,
	}
}

func (stc *SchedulerTriggerCreate) Execute(scheduleReq dto.SchedulerTriggerReq) (*string, error) {
	scheduleReq.UID = uuid.New()

	// TODO: use `repository` and `domain` to make it `Tell dont ask` in future

	ctxReq := context.Background()
	key := fmt.Sprintf("schedule:%s", scheduleReq.UID.String())

	// CACHED MESSAGE DATA
	expiration, err := stc.cacheInMemoDB.GetDefaultExpiration(ctxReq)
	if err != nil {
		return nil, fmt.Errorf("failed on cache GetDefaultExpiration: %w", err)
	}

	err = stc.cacheInMemoDB.Set(
		ctxReq,
		key,
		scheduleReq,
		expiration,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to persists cacheInMemoDB data: %w", err)
	}

	// ONLY KEY TO TRIGGER SEND MESSAGE
	triggerAt, err := mapUTCDataToTimeDuration(scheduleReq.TriggerAt)
	if err != nil {
		return nil, fmt.Errorf("failed to MapUTCDataToTimeDuration: %w", err)
	}

	err = stc.triggerInMemoDB.Set(
		ctxReq,
		key,
		nil,
		*triggerAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to persists triggerInMemoDB data: %w", err)
	}

	uid := scheduleReq.UID.String()
	return &uid, nil
}
