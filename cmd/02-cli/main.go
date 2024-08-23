package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-on-go",
	Short: "Go on Go is a Go lang workshop",
	Long:  `Go on Go is a Go lang workshop`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Go on Go",
	Long:  `Print the version number of Go on Go`,
	Run: func(cmd *cobra.Command, args []string) {
		println("Go on Go v1.0.0")
	},
}

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Print Hello, World!",
	Long:  `Print Hello, World!`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		println("Hello,", args[0])
	},
}

var repeatCmd = &cobra.Command{
	Use:   "repeat",
	Short: "Print a string multiple times",
	Long:  `Print a string multiple times`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repeat, err := cmd.Flags().GetInt("repeat")

		if err != nil {
			return err
		}

		for i := 0; i < repeat; i++ {
			println(args[0])
		}

		return nil
	},
}

func makeRepeatCmd() *cobra.Command {
	repeatCmd.Flags().IntP("repeat", "r", 1, "Number of times to repeat the string")

	return repeatCmd
}

func main() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(helloCmd)
	rootCmd.AddCommand(makeRepeatCmd())

	err := rootCmd.Execute()

	if err != nil {
		println(err.Error())
	}
}
