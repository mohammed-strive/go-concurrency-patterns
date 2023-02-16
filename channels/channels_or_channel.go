package channels

import (
	"fmt"
	"time"
)

func Or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-Or(append(channels[3:], orDone)...):
			}
		}
	}()

	return orDone
}

func OrChannelExample() {
	stg := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()

		return c
	}

	start := time.Now()
	<-Or(
		stg(2*time.Hour),
		stg(5*time.Second),
		stg(5*time.Minute),
		stg(1*time.Hour),
		stg(1*time.Minute),
	)

	fmt.Printf("done after %v", time.Since(start))
}
