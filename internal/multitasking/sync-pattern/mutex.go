package syncPattern

import (
	"context"
	"fmt"
	"sync"
)

type SharedResource struct {
	lock  sync.RWMutex // sync.Mutex
	value int
}

func NewSharedResource() *SharedResource {
	return &SharedResource{}
}

func (sharedResource *SharedResource) incrementUnsafe() {
	sharedResource.value++
}

func (sharedResource *SharedResource) Increment() {
	sharedResource.lock.Lock()
	defer sharedResource.lock.Unlock()

	sharedResource.incrementUnsafe()
}

func NoMutexShowcase(_ context.Context) {
	sharedResource := NewSharedResource()
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			sharedResource.incrementUnsafe()
		}()
	}

	wg.Wait()

	fmt.Println(sharedResource.value)
}

func WithMutexShowcase(_ context.Context) {
	sharedResource := NewSharedResource()

	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			sharedResource.Increment()
		}()
	}

	wg.Wait()

	fmt.Println(sharedResource.value)
}
