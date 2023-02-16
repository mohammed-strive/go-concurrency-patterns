package channels

func OrDone(done <-chan interface{}, input <-chan interface{}) <-chan interface{} {
	output := make(chan interface{})

	go func() {
		defer close(output)

		for {
			select {
			case <-done:
				return
			case v, ok := <-input:
				if !ok {
					return
				}
				select {
				case output <- v:
				case <-done:
					return
				}
			}
		}
	}()

	return output
}
