package main

import "goOnGo/internal/orders/use-case/cobra"

func main() {
	if err := cobra.RootCmd().Execute(); err != nil {
		panic(err)
	}
}
