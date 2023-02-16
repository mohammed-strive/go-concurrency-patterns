package channels

import "fmt"

func Take(done <-chan interface{}, values <-chan interface{}, num int) <-chan interface{} {
	stream := make(chan interface{})

	go func() {
		defer close(stream)

		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case stream <- <-values:
			}
		}
	}()

	return stream
}

func TakeExample() {
	done := make(chan interface{})
	take := Take(done, Repeat(done, 1, 2, 3, 4, 5), 7)

	for val := range take {
		fmt.Printf("Take Val: %v\n", val)
	}
}
