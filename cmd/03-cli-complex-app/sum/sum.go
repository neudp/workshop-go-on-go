package sum

import "github.com/spf13/cobra"

func New() *cobra.Command {
	var sumCmd = &cobra.Command{
		Use:   "sum",
		Short: "Sum two numbers",
	}

	sumCmd.AddCommand(newSumIntCmd())
	sumCmd.AddCommand(newSumFloatCmd())

	return sumCmd
}
