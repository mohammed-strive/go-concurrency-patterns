package channels

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	NUM_GOROUTINES uint64 = 10_000_000
)

func ChannelsGoroutineSize() {
	var wg sync.WaitGroup
	wg.Add(int(NUM_GOROUTINES))

	var stats runtime.MemStats

	block := make(chan interface{})
	close(block)

	for i := 0; i < int(NUM_GOROUTINES); i++ {
		go func() {
			wg.Done()
			<-block
		}()
	}

	wg.Wait()
	runtime.ReadMemStats(&stats)
	fmt.Printf("Memory Stats - memory consumed per go routine - %v", stats.Sys)
}
