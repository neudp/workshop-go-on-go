package sum

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func floatSum(numbers ...float64) float64 {
	var total float64
	for _, number := range numbers {
		total += number
	}
	return total
}

func newSumFloatCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "float",
		Short: "Sum float numbers",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			precision, err := cmd.Flags().GetInt("precision")

			if err != nil {
				return err
			}

			numbers := make([]float64, len(args))

			for index, input := range args {
				number, err := strconv.ParseFloat(input, 64)

				if err != nil {
					return err
				}

				numbers[index] = number
			}

			result := floatSum(numbers...)

			fmt.Printf("%0.*f\n", precision, result)

			return nil
		},
	}

	cmd.Flags().IntP("precision", "p", 2, "Number of decimal places to round the result to")

	return cmd
}
