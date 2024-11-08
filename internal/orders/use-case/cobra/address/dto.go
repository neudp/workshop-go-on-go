package address

import "goOnGo/internal/orders/model/domain/address"

type CountryDto struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

func FromCountry(country *address.Country) *CountryDto {
	return &CountryDto{
		Id:   uint64(country.Id()),
		Name: country.Name(),
	}
}

type ProvinceDto struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

func FromProvince(province *address.Province) *ProvinceDto {
	return &ProvinceDto{
		Id:   uint64(province.Id()),
		Name: province.Name(),
	}
}

type CityDto struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

func FromCity(city *address.City) *CityDto {
	return &CityDto{
		Id:   uint64(city.Id()),
		Name: city.Name(),
	}
}

type AddressDto struct {
	Street      string       `json:"street"`
	PostalCode  string       `json:"postal_code"`
	AddressLine string       `json:"address_line"`
	Country     *CountryDto  `json:"country"`
	Province    *ProvinceDto `json:"province"`
	City        *CityDto     `json:"city"`
}

func FromAddress(addr *address.Address) *AddressDto {
	return &AddressDto{
		Street:      string(addr.Street()),
		PostalCode:  string(addr.PostalCode()),
		AddressLine: string(addr.AddressLine()),
		Country:     FromCountry(addr.Country()),
		Province:    FromProvince(addr.Province()),
		City:        FromCity(addr.City()),
	}
}
