package syncPattern

import (
	"fmt"
	"sync"
	"time"
)

func Pipeline(name string, cond *sync.Cond, wg *sync.WaitGroup) {
	defer wg.Done()

	// sync short time part
	cond.L.Lock() // wait for main to start
	cond.L.Unlock()

	// async long time part
	<-time.After(time.Second) // Do some long work

	fmt.Printf("%s end phase 1\n", name)

	// sync short time part
	cond.L.Lock()
	wg.Done()
	cond.Wait() // Wait until all goroutines are done with phase 1
	cond.L.Unlock()

	// async long time part
	<-time.After(time.Second) // Do some long work

	fmt.Printf("%s end phase 2\n", name)
}

func CondShowCase() {
	workers := 1000
	start := time.Now()
	cond := sync.NewCond(&sync.Mutex{})

	wg := sync.WaitGroup{}

	for i := 0; i < workers; i++ {
		go Pipeline(fmt.Sprintf("Pipeline %d", i), cond, &wg)
	}

	fmt.Println("Start phase 1")
	wg.Add(workers)
	cond.Broadcast() // Wake up all goroutines

	wg.Wait()

	fmt.Println("Start phase 2")
	wg.Add(workers)
	cond.Broadcast()

	wg.Wait()

	fmt.Printf("Elapsed time: %s\n", time.Since(start))
}
