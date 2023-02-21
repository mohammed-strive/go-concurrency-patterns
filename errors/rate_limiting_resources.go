package errors

import (
	"context"
	"time"

	"golang.org/x/time/rate"
)

type APIResourceConnection struct {
	networkLimit,
	diskLimit,
	apiLimit RateLimiter
}

func (a *APIResourceConnection) ReadFile(ctx context.Context) error {
	err := MultiLimiter(a.apiLimit, a.diskLimit).Wait(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (a *APIResourceConnection) ResolveAddress(ctx context.Context) error {
	err := MultiLimiter(a.apiLimit, a.networkLimit).Wait(ctx)
	if err != nil {
		return err
	}

	return nil
}

func OpenResource() *APIResourceConnection {
	return &APIResourceConnection{
		apiLimit: MultiLimiter(
			rate.NewLimiter(Per(2, time.Second), 2),
			rate.NewLimiter(Per(10, time.Minute), 10),
		),
		diskLimit: MultiLimiter(
			rate.NewLimiter(rate.Limit(1), 1),
		),
		networkLimit: MultiLimiter(
			rate.NewLimiter(Per(3, time.Second), 3),
		),
	}
}

func MultiLimiterResource() {
	apiConnection := OpenResource()
	TestRateLimiting(apiConnection)
}
