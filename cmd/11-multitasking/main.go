package main

import (
	"context"
	"fmt"
	"goOnGo/internal/multitasking"
	syncPattern "goOnGo/internal/multitasking/sync-pattern"
)

var showcaseMap = map[string]func(context.Context){
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

var showcaseOrder = []string{
	"sync-pattern/no-wait-group",
	"sync-pattern/with-wait-group",
	"sync-pattern/no-mutex",
	"sync-pattern/with-mutex",
	"sync-pattern/cond",
	"sync-pattern/once-func",
	"sync-pattern/once-value",
	"parent-context",
	"generator",
	"worker-pool/round-robin",
	"worker-pool/least-connections",
}

func main() {
	fmt.Println("Hello, playground")

	cursor := 0
	for {
		key := showcaseOrder[cursor]
		fmt.Printf("You are here: %s\n", key)

		fmt.Println("Actions:")
		fmt.Println(">. Run")
		fmt.Println(">>. Next")
		fmt.Println("<<. Previous")
		fmt.Println("q. Quit")

		var action string
		if _, err := fmt.Scanln(&action); err != nil {
			fmt.Println("Invalid input")

			continue
		}

		switch action {
		case ">":
			fmt.Printf("Running %s\n", key)
			fmt.Println("========================================")
			ctx, cancel := context.WithCancel(context.Background())
			showcaseMap[key](ctx)
			cancel()
			fmt.Println("========================================")
		case ">>":
			cursor = (cursor + 1) % len(showcaseOrder)
		case "<<":
			cursor = (cursor - 1 + len(showcaseOrder)) % len(showcaseOrder)
		case "q":
			fmt.Println("Goodbye")
			return
		}
	}
}
