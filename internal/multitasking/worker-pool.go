package multitasking

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

func RoundRobin(ctx context.Context, concurrency int, tasks <-chan func(int)) {
	stack := make([]chan func(int), cap(tasks))
	for index := range stack {
		stack[index] = make(chan func(int), concurrency)
	}

	load := make([]int, cap(stack))
	lock := sync.Mutex{}
	for index, channel := range stack {
		go func(index int, channel chan func(int)) {
			for {
				select {
				case <-ctx.Done():
					return
				case task := <-channel:
					task(index)

					lock.Lock()
					load[index]--
					lock.Unlock()
				}
			}
		}(index, channel)
	}

	go func() {
		elected := 0
		for {
			select {
			case <-ctx.Done():
				return
			case task, open := <-tasks:
				if !open && task == nil {
					return
				}

				for load[elected] == concurrency { // full
					elected = (elected + 1) % concurrency
				}

				lock.Lock()
				load[elected]++
				lock.Unlock()

				stack[elected] <- task

				elected = (elected + 1) % concurrency
			}
		}
	}()
}

func LeastConnections(ctx context.Context, concurrency int, tasks <-chan func(int)) {
	stack := make([]chan func(int), cap(tasks))
	for index := range stack {
		stack[index] = make(chan func(int), concurrency)
	}

	load := make([]int, cap(stack))
	lock := sync.Mutex{}
	for index, channel := range stack {
		go func(index int, channel chan func(int)) {
			for {
				select {
				case <-ctx.Done():
					return
				case task := <-channel:
					task(index)

					lock.Lock()
					load[index]--
					lock.Unlock()
				}
			}
		}(index, channel)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case task, open := <-tasks:
				if !open && task == nil {
					return
				}

				elected := 0
				electedLoad := load[elected]
				for candidate, candidateLoad := range load {
					if candidateLoad < electedLoad {
						elected = candidate
						electedLoad = candidateLoad
					}
				}

				lock.Lock()
				load[elected]++
				lock.Unlock()

				stack[elected] <- task
			}
		}
	}()
}

func RoundRobinShowcase(ctx context.Context) {
	count := 1000
	tasks := make(chan func(int), 10)
	ctx, cancel := context.WithCancel(ctx)
	doerLoad := make(map[int]int)
	doerTime := make(map[int]time.Duration)
	taskStart := make(map[int]time.Time)
	taskWait := make([]time.Duration, count)
	lock := sync.Mutex{}

	var wg sync.WaitGroup
	wg.Add(count)
	go func() {
		for i := 0; i < count; i++ {
			taskStart[i] = time.Now()
			tasks <- func(doer int) {
				defer wg.Done()

				lock.Lock()
				taskWait[i] = time.Since(taskStart[i])
				doerLoad[doer]++
				lock.Unlock()

				start := time.Now()
				<-time.After(time.Duration(rand.Intn(30)+10) * time.Millisecond)
				<-time.After(10 * time.Millisecond)

				lock.Lock()
				doerTime[doer] += time.Since(start)
				lock.Unlock()
			}
			<-time.After(time.Duration(rand.Intn(10)+10) * time.Millisecond)
		}
	}()

	RoundRobin(ctx, 10, tasks)

	wg.Wait()
	cancel()

	fmt.Printf("Doer load: %v\n", doerLoad)
	fmt.Printf("Doer time: %v\n", doerTime)

	squareSum := 0.0
	for _, value := range taskWait {
		squareSum += float64(value.Milliseconds()) * float64(value.Milliseconds())
	}

	fmt.Printf("Task wait variance: %v\n", math.Pow(squareSum/float64(count), 0.5))
}

func LeastConnectionsShowcase(ctx context.Context) {
	count := 1000
	tasks := make(chan func(int), 10)
	ctx, cancel := context.WithCancel(ctx)
	doerLoad := make(map[int]int)
	doerTime := make(map[int]time.Duration)
	taskStart := make(map[int]time.Time)
	taskWait := make([]time.Duration, count)
	lock := sync.Mutex{}

	var wg sync.WaitGroup
	wg.Add(count)
	go func() {
		for i := 0; i < count; i++ {
			taskStart[i] = time.Now()
			tasks <- func(doer int) {
				defer wg.Done()

				lock.Lock()
				taskWait[i] = time.Since(taskStart[i])
				doerLoad[doer]++
				lock.Unlock()

				start := time.Now()
				<-time.After(time.Duration(rand.Intn(30)+10) * time.Millisecond)
				<-time.After(10 * time.Millisecond)

				lock.Lock()
				doerTime[doer] += time.Since(start)
				lock.Unlock()
			}
			<-time.After(time.Duration(rand.Intn(10)+10) * time.Millisecond)
		}
	}()

	LeastConnections(ctx, 10, tasks)

	wg.Wait()
	cancel()

	fmt.Printf("Doer load: %v\n", doerLoad)
	fmt.Printf("Doer time: %v\n", doerTime)

	squareSum := 0.0
	for _, value := range taskWait {
		squareSum += float64(value.Milliseconds()) * float64(value.Milliseconds())
	}

	fmt.Printf("Task wait variance: %v\n", math.Pow(squareSum/float64(count), 0.5))
}
