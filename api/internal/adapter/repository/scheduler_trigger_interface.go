package repository

import (
	"context"
	"time"

	"github.com/jtonynet/go-scheduler-trigger/api/internal/core/dto"
)

type SchedulerTrigger interface {
	Create(ctx context.Context, key string, scheduleReq *dto.SchedulerTriggerReq, expiration *time.Duration) error
}
