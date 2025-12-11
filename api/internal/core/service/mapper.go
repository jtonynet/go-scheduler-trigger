package service

import (
	"errors"
	"fmt"
	"time"
)

func mapUTCDataToTimeDuration(UTCData string) (*time.Duration, error) {
	targetTime, err := time.Parse(time.RFC3339, UTCData)
	if err != nil {
		return nil, fmt.Errorf("UTC parser error: %w", err)
	}

	now := time.Now().UTC()
	duration := targetTime.Sub(now)

	if duration <= 0 {
		return nil, errors.New("UTC data bellow actual time")
	}

	return &duration, nil
}
