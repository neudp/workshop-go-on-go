package syncPattern

import (
	"fmt"
	"sync"
)

func NoWaitGroupShowcase() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("Hello from worker %d\n", i)
		}()
	}
}

func WithWaitGroupShowcase() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			fmt.Printf("Hello from worker %d\n", i)
		}()
	}

	wg.Wait()
}
