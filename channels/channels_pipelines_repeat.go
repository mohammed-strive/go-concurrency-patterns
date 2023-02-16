package channels

import (
	"fmt"
	"time"
)

func RepeatFn(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	stream := make(chan interface{})

	go func() {
		defer close(stream)

		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()

	return stream
}

func Repeat(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	stream := make(chan interface{})

	go func() {
		defer close(stream)

		for {
			for _, val := range values {
				select {
				case <-done:
					return
				case stream <- val:
				}
			}
		}
	}()

	return stream
}

func RepeatExample() {
	done := make(chan interface{})
	defer close(done)
	values := Repeat(done, 1, 2, 3, 4, 5)

	for {
		select {
		case <-done:
			return
		case val := <-values:
			fmt.Printf("Values: %v\n", val)
			time.Sleep(1 * time.Second)
		case <-time.After(2 * time.Second):
			return
		}
	}
}
