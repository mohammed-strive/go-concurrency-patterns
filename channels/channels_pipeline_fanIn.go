package channels

import "sync"

func FanIn(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
	fanInStream := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(len(channels))

	go func() {
		for _, channel := range channels {
			channel := channel
			go func() {
				defer wg.Done()
				for i := range channel {
					select {
					case <-done:
						return
					case fanInStream <- i:
					}
				}
			}()
		}
	}()

	go func() {
		wg.Wait()
		close(fanInStream)
	}()

	return fanInStream
}
