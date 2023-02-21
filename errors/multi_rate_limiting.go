package errors

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
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

func TestRateLimiting(api ApiConnectionInterface) {
	defer fmt.Println("Done")

	log.SetFlags(log.LUTC | log.Ltime)
	log.SetOutput(os.Stdout)

	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			if err := api.ReadFile(context.Background()); err != nil {
				log.Printf("cannot read file: %v", err)
			}
			log.Println("read file")
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			if err := api.ResolveAddress(context.Background()); err != nil {
				log.Printf("cannot resolve address: %v", err)
			}
			log.Println("resolve address")
		}()
	}

	wg.Wait()
}

func MultiRateLimiting() {
	defer fmt.Println("Done")

	log.SetOutput(os.Stdout)
	log.SetFlags(log.LUTC | log.Ltime)

	apiConnection := OpenForMultiRateLimiter()
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			if err := apiConnection.ReadFile(context.Background()); err != nil {
				log.Printf("cannot read file: %v", err)
			}
			log.Println("read file")
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			if err := apiConnection.ResolveAddress(context.Background()); err != nil {
				log.Printf("cannot resolve address: %v", err)
			}
			log.Println("resolve address")
		}()
	}

	wg.Wait()
}
