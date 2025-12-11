package repository

import (
	"context"
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

	err = str.db.Set(
		ctx,
		key,
		scheduleReq,
		*expiration,
	)
	if err != nil {
		return fmt.Errorf("failed to create SchedulerTriggerRedis: %w", err)
	}

	return nil
}
