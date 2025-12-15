package bootstrap

import (
	"fmt"

	"github.com/jtonynet/go-scheduler-trigger/api/config"
	"github.com/jtonynet/go-scheduler-trigger/api/internal/adapter/database"
	"github.com/jtonynet/go-scheduler-trigger/api/internal/adapter/email"
	"github.com/jtonynet/go-scheduler-trigger/api/internal/adapter/pubSub"
	"github.com/jtonynet/go-scheduler-trigger/api/internal/adapter/repository"
	"github.com/jtonynet/go-scheduler-trigger/api/internal/core/service"
)

type REST struct {
	SchedulerTriggerCreate *service.SchedulerTriggerCreate
}

type Worker struct {
	TriggerPubSub           pubSub.PubSub
	SchedulerTriggerExpired *service.SchedulerTriggerExpired
}

func NewREST(cfg *config.Config) (*REST, error) {
	shadowKeyDB, err := database.NewInMemory(cfg.ShadowKeyDB.ToInMemoryDB())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize shadowKeyDB: %w", err)
	}

	triggerDB, err := database.NewInMemory(cfg.TriggerDB.ToInMemoryDB())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize triggerDB: %w", err)
	}

	shadowKeyRepo := repository.NewSchedulerTriggerRedis(shadowKeyDB)
	triggerRepo := repository.NewSchedulerTriggerRedis(triggerDB)

	schedulerTriggerCreate := service.NewSchedulerTriggerCreate(
		shadowKeyRepo,
		triggerRepo,
	)

	return &REST{
		SchedulerTriggerCreate: schedulerTriggerCreate,
	}, nil
}

func NewWorker(cfg *config.Config) (*Worker, error) {
	mail := email.New(cfg.MailNotification)

	triggerPubSub, err := pubSub.NewRedisPubSub(cfg.PubSub)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize triggerPubSub: %w", err)
	}

	shadowKeyDB, err := database.NewInMemory(cfg.ShadowKeyDB.ToInMemoryDB())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize shadowKeyDB: %w", err)
	}
	shadowKeyRepo := repository.NewSchedulerTriggerRedis(shadowKeyDB)

	schedulerTriggerExpired := service.NewSchedulerTriggerExpired(
		*mail,
		shadowKeyRepo,
	)

	return &Worker{
		TriggerPubSub:           triggerPubSub,
		SchedulerTriggerExpired: schedulerTriggerExpired,
	}, nil
}
