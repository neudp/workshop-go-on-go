package address

type CityId uint64

type City struct {
	id       CityId
	country  *Country
	province *Province
	name     string
}

func NewCity(name string, country *Country, province *Province) *City {
	return &City{
		country:  country,
		province: province,
		name:     name,
	}
}

func RestoreCity(id CityId, name string, country *Country, province *Province) *City {
	return &City{
		id:       id,
		country:  country,
		province: province,
		name:     name,
	}
}

func (city *City) Id() CityId {
	return city.id
}

func (city *City) Country() *Country {
	return city.country
}

func (city *City) Province() *Province {
	return city.province
}

func (city *City) Name() string {
	return city.name
}

func (city *City) Clone() *City {
	newCity := new(City)
	*newCity = *city

	return newCity
}

func (city *City) setCountry(country *Country) {
	city.country = country
}

func (city *City) setProvince(province *Province) {
	city.province = province
}

func (city *City) WithId(id CityId) *City {
	newCity := city.Clone()
	newCity.id = id

	return newCity
}

func (city *City) Rename(name string) *City {
	newCity := city.Clone()
	newCity.name = name

	return newCity
}

func (city *City) ChangeCountry(country *Country) *City {
	newCity := city.Clone()
	newCity.setCountry(country)
	newCity.province = newCity.province.Clone()
	newCity.province.setCountry(country)

	return newCity
}

func (city *City) ChangeProvince(province *Province) *City {
	newCity := city.Clone()
	newCity.province = province
	newCity.country = province.Country()

	return newCity
}
