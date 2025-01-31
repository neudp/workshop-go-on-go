package multitasking

import (
	"context"
	"fmt"
	"time"
)

func IntGenerator(start, end int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := start; i <= end; i++ {
			ch <- i
			time.Sleep(100 * time.Millisecond) // simulate work
		}
		close(ch)
	}()
	return ch
}

func GeneratorShowcase(_ context.Context) {
	gen := IntGenerator(1, 10)
	for num := range gen {
		fmt.Printf("Generated number: %d\n", num)
	}
}
