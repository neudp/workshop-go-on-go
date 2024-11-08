package address

type ProvinceId uint64

type Province struct {
	id      ProvinceId
	name    string
	country *Country
}

func NewProvince(name string, country *Country) *Province {
	return &Province{name: name, country: country}
}

func RestoreProvince(id ProvinceId, name string, country *Country) *Province {
	return &Province{id: id, name: name, country: country}
}

func (province *Province) Id() ProvinceId {
	return province.id
}

func (province *Province) Country() *Country {
	return province.country
}

func (province *Province) Name() string {
	return province.name
}

func (province *Province) Clone() *Province {
	newProvince := new(Province)
	*newProvince = *province

	return newProvince
}

func (province *Province) setCountry(country *Country) {
	province.country = country
}

func (province *Province) WithId(id ProvinceId) *Province {
	newProvince := province.Clone()
	newProvince.id = id

	return newProvince
}

func (province *Province) Rename(name string) *Province {
	newProvince := province.Clone()
	newProvince.name = name

	return newProvince
}

func (province *Province) ChangeCountry(country *Country) *Province {
	newProvince := province.Clone()
	newProvince.setCountry(country)

	return newProvince
}
