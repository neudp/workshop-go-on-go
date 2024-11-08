package address

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	addressCreate "goOnGo/internal/orders/application/address-management/address-create"
	addressRepository "goOnGo/internal/orders/infrastructure/repository/address-repository"
	"goOnGo/internal/orders/model/domain/address"
)

type Country struct {
	IdValue   address.CountryId `json:"id"`
	NameValue string            `json:"name"`
}

func (country *Country) Id() address.CountryId {
	return country.IdValue
}

func (country *Country) Name() string {
	return country.NameValue
}

type Province struct {
	IdValue   address.ProvinceId `json:"id"`
	NameValue string             `json:"name"`
}

func (province *Province) Id() address.ProvinceId {
	return province.IdValue
}

func (province *Province) Name() string {
	return province.NameValue
}

type City struct {
	IdValue   address.CityId `json:"id"`
	NameValue string         `json:"name"`
}

func (city *City) Id() address.CityId {
	return city.IdValue
}

func (city *City) Name() string {
	return city.NameValue
}

type CreateAddressRequest struct {
	StreetValue      address.Street      `json:"street"`
	PostalCodeValue  address.PostalCode  `json:"postal_code"`
	AddressLineValue address.AddressLine `json:"address_line"`
	CountryValue     *Country            `json:"country"`
	ProvinceValue    *Province           `json:"province"`
	CityValue        *City               `json:"city"`
}

func (request *CreateAddressRequest) Street() address.Street {
	return request.StreetValue
}

func (request *CreateAddressRequest) PostalCode() address.PostalCode {
	return request.PostalCodeValue
}

func (request *CreateAddressRequest) AddressLine() address.AddressLine {
	return request.AddressLineValue
}

func (request *CreateAddressRequest) Country() addressCreate.Country {
	return request.CountryValue
}

func (request *CreateAddressRequest) Province() addressCreate.Province {
	return request.ProvinceValue
}

func (request *CreateAddressRequest) City() addressCreate.City {
	return request.CityValue
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an address",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		payload := new(CreateAddressRequest)
		if err := json.Unmarshal([]byte(args[0]), payload); err != nil {
			return err
		}

		db, err := sql.Open("mysql", "go_on_go:go_on_go@tcp(127.0.0.1:8306)/go_on_go")
		if err != nil {
			return err
		}

		repository := addressRepository.New(db)
		createAddress := addressCreate.NewHandler(repository)

		addr, err := createAddress.Handle(payload)
		if err != nil {
			return err
		}

		responseBytes, err := json.MarshalIndent(FromAddress(addr), "", "  ")
		if err != nil {
			return err
		}

		cmd.Println(string(responseBytes))

		return nil
	},
}

func CreateCmd() *cobra.Command {
	return createCmd
}
