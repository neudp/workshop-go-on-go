package addressRepository

import (
	"database/sql"
	_ "embed"
	"errors"
	"goOnGo/internal/orders/model/domain/address"
)

var (
	//go:embed query/insert-address.sql
	insertAddressQuery string
	//go:embed query/select-address-all.sql
	selectAddressIdByNameCountryProvinceCityQuery string
)

type AddressRepository struct {
	db                 *sql.DB
	countryRepository  *CountryRepository
	provinceRepository *ProvinceRepository
	cityRepository     *CityRepository
}

func New(db *sql.DB) *AddressRepository {
	return &AddressRepository{db: db}
}

func (*AddressRepository) save(
	tx *sql.Tx,
	addr *address.Address,
) (addressId address.Id, err error) {
	stmt, err := tx.Exec(
		insertAddressQuery,
		addr.Street(),
		addr.PostalCode(),
		addr.AddressLine(),
		addr.Country().Id(),
		addr.Province().Id(),
		addr.City().Id(),
	)

	if err != nil {
		return 0, err
	}

	id, err := stmt.LastInsertId()
	if err != nil {
		return 0, err
	}

	return address.Id(id), err
}

func (*AddressRepository) findAll(tx *sql.Tx) (addrs []*address.Address, err error) {
	rows, err := tx.Query(selectAddressIdByNameCountryProvinceCityQuery)

	if err != nil {
		return nil, err
	}

	defer func() {
		err = errors.Join(err, rows.Close())
	}()

	countries := make(map[address.CountryId]*address.Country)
	provinces := make(map[address.ProvinceId]*address.Province)
	cities := make(map[address.CityId]*address.City)
	addrs = make([]*address.Address, 0)

	for rows.Next() {
		if err = rows.Err(); err != nil {
			return nil, err
		}

		var (
			id           address.Id
			street       address.Street
			postalCode   address.PostalCode
			addressLine  address.AddressLine
			countryId    address.CountryId
			countryName  string
			provinceId   address.ProvinceId
			provinceName string
			cityId       address.CityId
			cityName     string
		)

		if err = rows.Scan(
			&id,
			&street,
			&postalCode,
			&addressLine,
			&countryId,
			&countryName,
			&provinceId,
			&provinceName,
			&cityId,
			&cityName,
		); err != nil {
			return nil, err
		}

		country, ok := countries[countryId]
		if !ok {
			country = address.RestoreCountry(countryId, countryName)
			countries[countryId] = country
		}

		province, ok := provinces[provinceId]
		if !ok {
			province = address.RestoreProvince(provinceId, provinceName, country)
			provinces[provinceId] = province
		}

		city, ok := cities[cityId]
		if !ok {
			city = address.RestoreCity(cityId, cityName, country, province)
			cities[cityId] = city
		}

		addrs = append(addrs, address.Restore(
			id,
			street,
			postalCode,
			addressLine,
			country,
			province,
			city,
		))
	}

	return addrs, nil
}

func (repository *AddressRepository) Save(addr *address.Address) (*address.Address, error) {
	tx, err := repository.db.Begin()

	if err != nil {
		return nil, err
	}

	if addr.Country().Id() == 0 {
		var countryId address.CountryId
		countryId, err = repository.countryRepository.findIdByName(tx, addr.Country().Name())
		if err != nil {
			countryId, err = repository.countryRepository.save(tx, addr.Country())
			if err != nil {
				return nil, errors.Join(err, tx.Rollback())
			}
		}

		addr = addr.ChangeCountry(addr.Country().WithId(countryId))
	}

	if addr.Province().Id() == 0 {
		var provinceId address.ProvinceId
		provinceId, err = repository.provinceRepository.findIdByNameCountry(
			tx,
			addr.Province().Name(),
			addr.Country().Id(),
		)
		if err != nil {
			provinceId, err = repository.provinceRepository.save(tx, addr.Province())
			if err != nil {
				return nil, errors.Join(err, tx.Rollback())
			}
		}

		addr = addr.ChangeProvince(addr.Province().WithId(provinceId))
	}

	if addr.City().Id() == 0 {
		var cityId address.CityId
		cityId, err = repository.cityRepository.findIdByNameCountryProvince(
			tx,
			addr.City().Name(),
			addr.Country().Id(),
			addr.Province().Id(),
		)
		if err != nil {
			cityId, err = repository.cityRepository.save(tx, addr.City())
			if err != nil {
				return nil, errors.Join(err, tx.Rollback())
			}
		}

		addr = addr.ChangeCity(addr.City().WithId(cityId))
	}

	addressId, err := repository.save(tx, addr)
	if err != nil {
		return nil, errors.Join(err, tx.Rollback())
	}

	return addr.WithId(addressId), tx.Commit()
}

func (repository *AddressRepository) FindAll() ([]*address.Address, error) {
	tx, err := repository.db.Begin()

	if err != nil {
		return nil, err
	}

	return repository.findAll(tx)
}
