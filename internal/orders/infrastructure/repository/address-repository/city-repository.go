package addressRepository

import (
	"database/sql"
	_ "embed"
	"errors"
	"goOnGo/internal/orders/model/domain/address"
)

var (
	//go:embed query/insert-city.sql
	insertCityQuery string
	//go:embed query/select-city-id-by-name-country-province.sql
	selectCityIdByNameCountryProvinceQuery string
)

type CityRepository struct {
	db                 *sql.DB
	countryRepository  *CountryRepository
	provinceRepository *ProvinceRepository
}

func NewCityRepository(
	db *sql.DB,
	countryRepository *CountryRepository,
	provinceRepository *ProvinceRepository,
) *CityRepository {
	return &CityRepository{
		db:                 db,
		countryRepository:  countryRepository,
		provinceRepository: provinceRepository,
	}
}

func (*CityRepository) save(
	tx *sql.Tx,
	city *address.City,
) (address.CityId, error) {
	stmt, err := tx.Exec(insertCityQuery, city.Name(), city.Country().Id(), city.Province().Id())

	if err != nil {
		return 0, err
	}

	id, err := stmt.LastInsertId()
	if err != nil {
		return 0, err
	}

	return address.CityId(id), err
}

func (*CityRepository) findIdByNameCountryProvince(
	tx *sql.Tx,
	cityName string,
	countryId address.CountryId,
	provinceId address.ProvinceId,
) (address.CityId, error) {
	var id int64
	err := tx.QueryRow(selectCityIdByNameCountryProvinceQuery, cityName, countryId, provinceId).Scan(&id)

	if err != nil {
		return 0, err
	}

	return address.CityId(id), nil
}

func (repository *CityRepository) Save(city *address.City) (*address.City, error) {
	tx, err := repository.db.Begin()
	if err != nil {
		return nil, err
	}

	if city.Country().Id() == 0 {
		var countryId address.CountryId
		countryId, err = repository.countryRepository.findIdByName(tx, city.Country().Name())
		if err != nil {
			countryId, err = repository.countryRepository.save(tx, city.Country())

			if err != nil {
				return nil, errors.Join(err, tx.Rollback())
			}
		}

		city = city.ChangeCountry(city.Country().WithId(countryId))
	}

	if city.Province().Id() == 0 {
		var provinceId address.ProvinceId
		provinceId, err = repository.provinceRepository.findIdByNameCountry(
			tx, city.Province().Name(),
			city.Country().Id(),
		)
		if err != nil {
			provinceId, err = repository.provinceRepository.save(tx, city.Province())

			if err != nil {
				return nil, errors.Join(err, tx.Rollback())
			}
		}

		city = city.ChangeProvince(city.Province().WithId(provinceId))
	}

	id, err := repository.save(tx, city)
	if err != nil {
		return nil, errors.Join(err, tx.Rollback())
	}

	return city.WithId(id), tx.Commit()
}
