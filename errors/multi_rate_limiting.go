package errors

import (
	"context"
	"time"

	"golang.org/x/time/rate"
)

type APIMultiConnection struct {
	rateLimiter RateLimiter
}

func (a *APIMultiConnection) ReadFile(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err
	}

	return nil
}

func (a *APIMultiConnection) ResolveAddress(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err
	}

	return nil
}

func Per(eventCount int, duration time.Duration) rate.Limit {
	return rate.Every(duration / time.Duration(eventCount))
}

func OpenForMultiRateLimiter() *APIMultiConnection {
	secondLimit := rate.NewLimiter(Per(2, time.Second), 1)
	minuteLimit := rate.NewLimiter(Per(10, time.Minute), 10)

	return &APIMultiConnection{
		rateLimiter: MultiLimiter(secondLimit, minuteLimit),
	}
}
