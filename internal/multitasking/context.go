package multitasking

import (
	"context"
	"fmt"
	"time"
)

func routine(ctx context.Context, ch <-chan int, name string) {
	for {
		select { // прерывание
		case <-ctx.Done():
			// выход
			return
		case i := <-ch:
			// задача
			fmt.Printf("Routine %s: %d\n", name, i)
		}
	}
}

func DoWithContext(ctx context.Context, name string) {
	ch := make(chan int)

	go routine(ctx, ch, name) // рутина != задача (процесс), рутина == неблокирующий шаг

	for i := 0; i < 10; i++ {
		ch <- i

		<-time.After(500 * time.Millisecond)
	}
}

func ParentContextShowcase(ctx context.Context) {
	subCtx1, cancel1 := context.WithCancel(ctx)
	subCtx2, cancel2 := context.WithCancel(ctx)

	fmt.Println("Child contexts")
	fmt.Println("start routines")

	go DoWithContext(subCtx1, "routine 1")
	go DoWithContext(subCtx2, "routine 2") // рутина 2

	fmt.Println("Stop routine 1")
	cancel1()

	<-time.After(1 * time.Second)

	fmt.Println("Stop routine 2")
	cancel2()

	fmt.Println("Parent context")
	fmt.Println("start routines")

	subCtx1, _ = context.WithTimeout(ctx, 10*time.Second)
	subCtx2, _ = context.WithTimeout(ctx, 10*time.Second)

	go DoWithContext(subCtx1, "routine 1")
	go DoWithContext(subCtx2, "routine 2")

	fmt.Println("Stop all routines")
}
