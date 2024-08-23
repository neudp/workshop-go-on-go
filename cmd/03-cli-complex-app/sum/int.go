package sum

import (
	"strconv"

	"github.com/spf13/cobra"
)

func intSum(numbers ...int64) int64 {
	var total int64
	for _, number := range numbers {
		total += number
	}

	return total
}

func newSumIntCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "int",
		Short: "Sum two numbers",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			numbers := make([]int64, len(args))

			for index, input := range args {
				number, err := strconv.Atoi(input)

				if err != nil {
					return err
				}

				numbers[index] = int64(number)
			}

			result := intSum(numbers...)

			println(result)

			return nil
		},
	}
}
