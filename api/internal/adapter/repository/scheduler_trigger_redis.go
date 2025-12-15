package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jtonynet/go-scheduler-trigger/api/internal/adapter/database"
	"github.com/jtonynet/go-scheduler-trigger/api/internal/core/dto"
)

type SchedulerTriggerRedis struct {
	db database.InMemory
}

func NewSchedulerTriggerRedis(db database.InMemory) SchedulerTrigger {
	return &SchedulerTriggerRedis{
		db,
	}
}

func (str *SchedulerTriggerRedis) Create(ctx context.Context, key string, scheduleReq *dto.SchedulerTriggerReq, expiration *time.Duration) error {
	var err error

	if expiration == nil {
		exp, err := str.db.GetDefaultExpiration(ctx)
		if err != nil {
			exp = 0
		}

		expiration = &exp
	}

	if err = str.db.Set(
		ctx,
		key,
		scheduleReq,
		*expiration,
	); err != nil {
		return fmt.Errorf("failed to create SchedulerTriggerRedis: %w", err)
	}

	return nil
}

func (str *SchedulerTriggerRedis) Retrieve(ctx context.Context, key string) (*dto.SchedulerTriggerReq, error) {
	st, err := str.db.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	var req dto.SchedulerTriggerReq
	if err := json.Unmarshal([]byte(st), &req); err != nil {
		return nil, fmt.Errorf("failed to unmarshal scheduler trigger: %w", err)
	}

	return &req, nil
}

func (str *SchedulerTriggerRedis) Delete(ctx context.Context, key string) error {
	if err := str.db.Delete(ctx, key); err != nil {
		return err
	}

	return nil
}
