package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jtonynet/go-scheduler-trigger/api/config"
)

type InMemory interface {
	Readiness(ctx context.Context) error
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
	Expire(ctx context.Context, key string, expiration time.Duration) error
	GetStrategy(ctx context.Context) (string, error)
	GetDefaultExpiration(ctx context.Context) (time.Duration, error)
	GetClient(ctx context.Context) (interface{}, error)
}

func NewInMemory(cfg config.InMemoryDatabase) (InMemory, error) {
	switch cfg.Strategy {
	case "redis":
		return NewRedisClient(cfg)
	default:
		return nil, fmt.Errorf("InMemoryDB strategy not suported: %s", cfg.Strategy)
	}
}
