package address

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "address",
	Short: "Address management",
}

func RootCmd() *cobra.Command {
	rootCmd.AddCommand(
		CreateCmd(),
		ListCmd(),
	)
	return rootCmd
}
