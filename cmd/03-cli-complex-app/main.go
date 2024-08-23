package main

import (
	"fmt"
	"goOnGo/cmd/03-cli-complex-app/sum"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := cobra.Command{
		Use:   "03-cli-complex-app",
		Short: "03-cli-complex-app is a complex CLI app example",
	}

	rootCmd.AddCommand(sum.New())

	err := rootCmd.Execute()

	if err != nil {
		fmt.Println(err.Error())
	}
}
