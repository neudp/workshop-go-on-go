package addressCreate

import "goOnGo/internal/orders/model/domain/address"

type Country interface {
	Id() address.CountryId
	Name() string
}

type Province interface {
	Id() address.ProvinceId
	Name() string
}

type City interface {
	Id() address.CityId
	Name() string
}

type CrateAddressRequest interface {
	Street() address.Street
	PostalCode() address.PostalCode
	AddressLine() address.AddressLine
	Country() Country
	Province() Province
	City() City
}

type Repository interface {
	Save(address *address.Address) (*address.Address, error)
}

type Handler struct {
	repository Repository
}

func NewHandler(repository Repository) *Handler {
	return &Handler{
		repository: repository,
	}
}

func (handler *Handler) Handle(request CrateAddressRequest) (*address.Address, error) {
	country := address.RestoreCountry(request.Country().Id(), request.Country().Name())
	province := address.RestoreProvince(request.Province().Id(), request.Province().Name(), country)
	city := address.RestoreCity(request.City().Id(), request.City().Name(), country, province)

	addr := address.New(
		request.Street(),
		request.PostalCode(),
		request.AddressLine(),
		country,
		province,
		city,
	)

	return handler.repository.Save(addr)
}
