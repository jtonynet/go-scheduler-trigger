package pubSub

import (
	"context"
	"fmt"

	"github.com/jtonynet/go-scheduler-trigger/api/config"
	"github.com/redis/go-redis/v9"
)

type RedisPubSub struct {
	client   *redis.Client
	pubsub   *redis.PubSub
	db       int
	strategy string
}

func NewRedisPubSub(cfg config.PubSub) (*RedisPubSub, error) {
	strAddr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	client := redis.NewClient(&redis.Options{
		Addr:     strAddr,
		Password: cfg.Pass,
		DB:       cfg.DB,
		Protocol: cfg.Protocol,
	})

	rps := &RedisPubSub{
		client:   client,
		db:       cfg.DB,
		strategy: cfg.Strategy,
	}

	return rps, nil
}

func (r *RedisPubSub) Subscribe(ctx context.Context) (<-chan string, error) {
	channel := fmt.Sprintf("__keyevent@%v__:expired", r.db)
	r.pubsub = r.client.Subscribe(ctx, channel)

	_, err := r.pubsub.Receive(ctx)
	if err != nil {
		return nil, fmt.Errorf("subscription channel error: %w", err)
	}

	out := make(chan string)

	go func() {
		defer close(out)

		ch := r.pubsub.Channel()
		for {
			select {
			case <-ctx.Done():
				return
			case msg, ok := <-ch:
				if !ok {
					return
				}
				out <- msg.Payload
			}
		}
	}()

	return out, nil
}

func (r *RedisPubSub) Close() error {
	if r.pubsub == nil {
		return nil
	}
	return r.pubsub.Close()
}

func (r *RedisPubSub) GetStrategy(_ context.Context) (string, error) {
	return r.strategy, nil
}
