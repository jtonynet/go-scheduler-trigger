package bootstrap

import (
	"fmt"

	"github.com/jtonynet/go-scheduler-trigger/api/config"
	"github.com/jtonynet/go-scheduler-trigger/api/internal/adapter/database"
	"github.com/jtonynet/go-scheduler-trigger/api/internal/core/service"
)

type REST struct {
	SchedulerTriggerCreate *service.SchedulerTriggerCreate
}

func NewREST(cfg *config.Config) (*REST, error) {
	cache, err := database.NewInMemory(cfg.Cache.ToInMemoryDB())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize cacheInMemoDB: %w", err)
	}

	trigger, err := database.NewInMemory(cfg.Trigger.ToInMemoryDB())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize triggerInMemoDB: %w", err)
	}

	schedulerTriggerCreate := service.NewSchedulerTriggerCreate(cache, trigger)

	return &REST{
		SchedulerTriggerCreate: schedulerTriggerCreate,
	}, nil
}
