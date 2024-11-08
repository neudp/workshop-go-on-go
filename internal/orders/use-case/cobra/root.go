package cobra

import (
	"github.com/spf13/cobra"
	"goOnGo/internal/orders/use-case/cobra/address"
)

var rootCmd = &cobra.Command{
	Use:   "order-service",
	Short: "Order service management",
}

func RootCmd() *cobra.Command {
	rootCmd.AddCommand(
		address.RootCmd(),
	)

	return rootCmd
}
