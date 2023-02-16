package channels

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func ChannelContext() {
	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := printGreetingContext(ctx); err != nil {
			fmt.Printf("cannot print greeting: %v\n", err)
			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := printFarewellContext(ctx); err != nil {
			fmt.Printf("cannot print farewell: %v\n", err)
		}
	}()

	wg.Wait()
}

func printGreetingContext(ctx context.Context) error {
	greeting, err := genGreetingContext(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("%s world!\n", greeting)
	return nil
}

func genGreetingContext(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	switch local, err := localeContext(ctx); {
	case err != nil:
		return "", err
	case local == "EN/US":
		return "hello ", nil
	}

	return "", fmt.Errorf("unsupported locale")
}

func localeContext(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(2 * time.Second):
	}

	return "EN/US", nil
}

func printFarewellContext(ctx context.Context) error {
	farewell, err := genFarewellContext(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("%s, world!\n", farewell)
	return nil
}

func genFarewellContext(ctx context.Context) (string, error) {
	switch local, err := localeContext(ctx); {
	case err != nil:
		return "", err
	case local == "EN/US":
		return "Goodbye", nil
	}

	return "", fmt.Errorf("unsupported locale")
}
