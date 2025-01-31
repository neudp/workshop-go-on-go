package syncPattern

import (
	"bytes"
	"context"
	"fmt"
	"sync"
)

func PoolShowcase(_ context.Context) {
	kilobytePool := sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 0, 1024))
		},
	}

	kilobyte := kilobytePool.Get().(*bytes.Buffer) // get buffer from pool

	// Do something with kilobyte

	kilobyte.Reset()
	kilobytePool.Put(kilobyte) // return buffer to pool
}

func AtomicMapShowcase(_ context.Context) {
	kvStore := sync.Map{}

	kvStore.Store("key", "value")                        // store key value
	value, ok := kvStore.Load("key")                     // get value if exists, ok is true if key exists
	kvStore.Delete("key")                                // delete key
	value, loaded := kvStore.LoadOrStore("key", "value") // get value if exists, store value if not exists, loaded is true if key existed
	value, ok = kvStore.LoadAndDelete("key")             // get value and delete key, ok is true if key existed
	kvStore.Range(func(key, value interface{}) bool {    // iterate over all key values
		// do something with key value
		return true // return false to stop iteration
	})
	value, ok = kvStore.Swap("key", "value") // set value and return previous, ok is true if key existed

	fmt.Println(value, ok, loaded)
}
