package address

type Id uint64
type Street string
type PostalCode string
type AddressLine string

type Address struct {
	id          Id
	country     *Country
	province    *Province
	city        *City
	street      Street
	postalCode  PostalCode
	addressLine AddressLine
}

func New(
	street Street,
	postalCode PostalCode,
	addressLine AddressLine,
	country *Country,
	province *Province,
	city *City,
) *Address {
	return &Address{
		country:     country,
		province:    province,
		city:        city,
		street:      street,
		postalCode:  postalCode,
		addressLine: addressLine,
	}
}

func Restore(
	id Id,
	street Street,
	postalCode PostalCode,
	addressLine AddressLine,
	country *Country,
	province *Province,
	city *City,
) *Address {
	return &Address{
		id:          id,
		country:     country,
		province:    province,
		city:        city,
		street:      street,
		postalCode:  postalCode,
		addressLine: addressLine,
	}
}

func (address *Address) Id() Id {
	return address.id
}

func (address *Address) Country() *Country {
	return address.country
}

func (address *Address) Province() *Province {
	return address.province
}

func (address *Address) City() *City {
	return address.city
}

func (address *Address) Street() Street {
	return address.street
}

func (address *Address) PostalCode() PostalCode {
	return address.postalCode
}

func (address *Address) AddressLine() AddressLine {
	return address.addressLine
}

func (address *Address) Clone() *Address {
	newAddress := new(Address)
	*newAddress = *address

	return newAddress
}

func (address *Address) WithId(id Id) *Address {
	newAddress := address.Clone()
	newAddress.id = id

	return newAddress
}

func (address *Address) setCountry(country *Country) {
	address.country = country
}

func (address *Address) setProvince(province *Province) {
	address.province = province
}

func (address *Address) setCity(city *City) {
	address.city = city
}

func (address *Address) ChangeCountry(country *Country) *Address {
	newAddress := address.Clone()
	newAddress.country = country
	newAddress.province = newAddress.province.Clone()
	newAddress.province.setCountry(country)
	newAddress.city = newAddress.city.Clone()
	newAddress.city.setCountry(country)
	newAddress.city.setProvince(newAddress.province)

	return newAddress
}

func (address *Address) ChangeProvince(province *Province) *Address {
	newAddress := address.Clone()
	newAddress.province = province
	newAddress.country = province.Country()
	newAddress.city = newAddress.city.Clone()
	newAddress.city.setProvince(province)
	newAddress.city.setCountry(province.Country())

	return newAddress
}

func (address *Address) ChangeCity(city *City) *Address {
	newAddress := address.Clone()
	newAddress.city = city
	newAddress.province = city.Province()
	newAddress.country = city.Country()

	return newAddress
}

func (address *Address) ChangeStreet(street Street) *Address {
	newAddress := address.Clone()
	newAddress.street = street

	return newAddress
}

func (address *Address) ChangePostalCode(postalCode PostalCode) *Address {
	newAddress := address.Clone()
	newAddress.postalCode = postalCode

	return newAddress
}

func (address *Address) ChangeAddressLine(addressLine AddressLine) *Address {
	newAddress := address.Clone()
	newAddress.addressLine = addressLine

	return newAddress
}
