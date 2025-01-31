package syncPattern

import (
	"context"
	"fmt"
	"sync"
)

func NoWaitGroupShowcase(ctx context.Context) {
	for i := 0; i < 10; i++ {
		go func() {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Printf("Hello from worker %d\n", i)
			}
		}()
	}
}

func WithWaitGroupShowcase(ctx context.Context) {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			select {
			case <-ctx.Done():
				return
			default:
				fmt.Printf("Hello from worker %d\n", i)
			}
		}()
	}

	wg.Wait()
}
