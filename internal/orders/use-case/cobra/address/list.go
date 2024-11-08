package address

import (
	"database/sql"
	"encoding/json"
	"github.com/spf13/cobra"
	addressList "goOnGo/internal/orders/application/address-management/address-list"
	addressRepository "goOnGo/internal/orders/infrastructure/repository/address-repository"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all addresses",
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := sql.Open("mysql", "go_on_go:go_on_go@tcp(127.0.0.1:8306)/go_on_go")
		if err != nil {
			return err
		}

		repository := addressRepository.New(db)
		listAddress := addressList.NewHandler(repository)

		addrs, err := listAddress.Handle()
		if err != nil {
			return err
		}

		response := make([]*AddressDto, len(addrs))
		for index, addr := range addrs {
			response[index] = FromAddress(addr)
		}

		responseBytes, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			return err
		}

		cmd.Println(string(responseBytes))

		return nil
	},
}

func ListCmd() *cobra.Command {
	return listCmd
}
