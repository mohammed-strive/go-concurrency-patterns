package errors

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"golang.org/x/time/rate"
)

func Open() *APIConnection {
	return &APIConnection{
		rateLimiter: rate.NewLimiter(rate.Limit(1), 1),
	}
}

type ApiConnectionInterface interface {
	ReadFile(context.Context) error
	ResolveAddress(context.Context) error
}

type APIConnection struct {
	rateLimiter *rate.Limiter
}

func (api *APIConnection) ReadFile(ctx context.Context) error {
	if err := api.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	return nil
}

func (api *APIConnection) ResolveAddress(ctx context.Context) error {
	if err := api.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	return nil
}

func RateLimiting() {
	defer fmt.Println("done")

	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	apiConnection := Open()
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			err := apiConnection.ReadFile(context.Background())
			if err != nil {
				log.Printf("cannot readfile: %v", err)
			}
			log.Println("read file")
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			err := apiConnection.ResolveAddress(context.Background())
			if err != nil {
				log.Printf("cannot resolve address: %v", err)
			}
			log.Println("resolve address")
		}()
	}

	wg.Wait()
}
