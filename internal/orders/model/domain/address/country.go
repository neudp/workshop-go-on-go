package address

type CountryId uint64

type Country struct {
	id   CountryId
	name string
}

func NewCountry(name string) *Country {
	return &Country{name: name}
}

func RestoreCountry(id CountryId, name string) *Country {
	return &Country{id: id, name: name}
}

func (country *Country) Id() CountryId {
	return country.id
}

func (country *Country) Name() string {
	return country.name
}

func (country *Country) Clone() *Country {
	newCountry := new(Country)
	*newCountry = *country

	return newCountry
}

func (country *Country) WithId(id CountryId) *Country {
	newCountry := country.Clone()
	newCountry.id = id

	return newCountry
}

func (country *Country) Rename(name string) *Country {
	newCountry := country.Clone()
	newCountry.name = name

	return newCountry
}
