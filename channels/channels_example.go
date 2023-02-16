package channels

import (
	"fmt"
	"sync"
)

func ChannelExample() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		fmt.Println("Hello, World from Channels")
	}()

	wg.Wait()
	fmt.Println("Hello, World from main")
}
