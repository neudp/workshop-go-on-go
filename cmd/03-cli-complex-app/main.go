package main

import (
	"fmt"
	"goOnGo/cmd/03-cli-complex-app/sum"

	"github.com/spf13/cobra"
)

/*
Cobra достаточно гибкий инструмент и позволяет создавать сложные CLI приложения.
В данном примере создается корневая команда rootCmd, которая содержит команду sum,
импортированную из пакета sum. В свою очередь, команда sum содержит две дочерние команды:
sumInt и sumFloat
*/
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
