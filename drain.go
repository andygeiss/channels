package channels

import (
	"runtime"
	"sync"
)

// Drain receives each value from an input channel and does something with it.
func Drain[T any](in <-chan T, fn func(in T)) chan bool {
	out := make(chan T)
	done := make(chan bool)
	// create a goroutine for each available cpu.
	num := runtime.NumCPU()
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			for val := range in {
				// process each value
				fn(val)
			}
		}()
	}
	// block until every channel is done.
	go func() {
		wg.Wait()
		close(out)
		close(done)
	}()
	return done
}
