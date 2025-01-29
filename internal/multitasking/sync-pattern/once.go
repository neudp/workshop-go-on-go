package syncPattern

import (
	"fmt"
	"sync"
)

func OnceFuncShowCase() {
	doOnce := sync.OnceFunc(func() {
		println("Only once")
	})

	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			doOnce()
		}()
	}

	wg.Wait()
}

func OnceValueShowCase() {
	//getOnce := sync.OnceValues(func() (int, error) {
	getOnce := sync.OnceValue(func() int {
		fmt.Println("Only once")
		return 42
	})

	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			value := getOnce()

			fmt.Println(value)
		}()
	}

	wg.Wait()
}
