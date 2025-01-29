package main

import (
	"goOnGo/internal/multitasking"
	syncPattern "goOnGo/internal/multitasking/sync-pattern"
	"os"
)

var showcaseMap = map[string]func(){
	"sync-pattern/no-wait-group":    syncPattern.NoWaitGroupShowcase,
	"sync-pattern/with-wait-group":  syncPattern.WithWaitGroupShowcase,
	"sync-pattern/no-mutex":         syncPattern.NoMutexShowcase,
	"sync-pattern/with-mutex":       syncPattern.WithMutexShowcase,
	"sync-pattern/cond":             syncPattern.CondShowCase,
	"sync-pattern/once-func":        syncPattern.OnceFuncShowCase,
	"sync-pattern/once-value":       syncPattern.OnceValueShowCase,
	"parent-context":                multitasking.ParentContextShowcase,
	"generator":                     multitasking.GeneratorShowcase,
	"worker-pool/round-robin":       multitasking.RoundRobinShowcase,
	"worker-pool/least-connections": multitasking.LeastConnectionsShowcase,
}

func main() {
	if len(os.Args) != 2 {
		panic("Invalid number of arguments")
	}

	showcase, ok := showcaseMap[os.Args[1]]

	if !ok {
		panic("Showcase not found")
	}

	showcase()
}
