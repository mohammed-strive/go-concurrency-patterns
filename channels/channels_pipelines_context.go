package channels

import (
	"fmt"
	"sync"
	"time"
)

func ChannelsContext() {
	var wg sync.WaitGroup

	done := make(chan interface{})
	defer close(done)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printGreeting(done); err != nil {
			fmt.Printf("%v", err)
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printFarewell(done); err != nil {
			fmt.Printf("%v", err)
			return
		}
	}()

	fmt.Println("Waiting for goroutines to finish")
	wg.Wait()
	fmt.Println("Goroutines are done... exiting")
}

func printFarewell(done <-chan interface{}) error {
	farewell, err := genFarewell(done)
	if err != nil {
		return err
	}

	fmt.Printf("%s world\n", farewell)
	return nil
}

func genFarewell(done <-chan interface{}) (string, error) {
	switch local, err := locale(done); {
	case err != nil:
		return "", nil
	case local == "EN/US":
		return "hello", nil
	}

	return "", fmt.Errorf("unsupported locale")
}

func printGreeting(done <-chan interface{}) error {
	greeting, err := genGreeting(done)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", greeting)
	return nil
}

func genGreeting(done <-chan interface{}) (string, error) {
	switch local, err := locale(done); {
	case err != nil:
		return "", err
	case local == "EN/US":
		return "goodbye", nil
	}

	return "", fmt.Errorf("unsupported locale")
}

func locale(done <-chan interface{}) (string, error) {
	select {
	case <-done:
		return "", fmt.Errorf("cancelled")
	case <-time.After(10 * time.Second):
	}

	return "EN/US", nil
}
