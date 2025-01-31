package multitasking

import (
	"context"
	"fmt"
	"time"
)

func IntGenerator(start, end int) <-chan int {
	ch := make(chan int, 10)
	go func() {
		for i := start; i <= end; i++ {
			ch <- i
			fmt.Printf("Generated number: %d\n", i)
		}
		close(ch)
	}()
	return ch
}

func GeneratorShowcase(_ context.Context) {
	gen := IntGenerator(1, 1000)

	for num := range gen {
		time.Sleep(100 * time.Millisecond) // simulate work
		fmt.Printf("Processed number: %d\n", num)
	}
}
