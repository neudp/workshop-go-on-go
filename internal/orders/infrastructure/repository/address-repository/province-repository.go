package addressRepository

import (
	"database/sql"
	_ "embed"
	"errors"
	"goOnGo/internal/orders/model/domain/address"
)

var (
	//go:embed query/insert-province.sql
	insertProvinceQuery string
	//go:embed query/select-province-id-by-name-country.sql
	selectProvinceIdByNameCountryQuery string
)

type ProvinceRepository struct {
	db                *sql.DB
	countryRepository *CountryRepository
}

func NewProvinceRepository(db *sql.DB, countryRepository *CountryRepository) *ProvinceRepository {
	return &ProvinceRepository{
		db:                db,
		countryRepository: countryRepository,
	}
}

func (*ProvinceRepository) save(
	tx *sql.Tx,
	province *address.Province,
) (address.ProvinceId, error) {
	stmt, err := tx.Exec(insertProvinceQuery, province.Name(), province.Country().Id())

	if err != nil {
		return 0, err
	}

	id, err := stmt.LastInsertId()
	if err != nil {
		return 0, err
	}

	return address.ProvinceId(id), err
}

func (*ProvinceRepository) findIdByNameCountry(
	tx *sql.Tx,
	provinceName string,
	countryId address.CountryId,
) (address.ProvinceId, error) {
	var id int64
	err := tx.QueryRow(selectProvinceIdByNameCountryQuery, provinceName, countryId).Scan(&id)

	if err != nil {
		return 0, err
	}

	return address.ProvinceId(id), nil
}

func (repository *ProvinceRepository) Save(province *address.Province) (*address.Province, error) {
	tx, err := repository.db.Begin()
	if err != nil {
		return nil, err
	}

	if province.Country().Id() == 0 {
		var countryId address.CountryId
		countryId, err = repository.countryRepository.findIdByName(tx, province.Country().Name())
		if err != nil {
			countryId, err = repository.countryRepository.save(tx, province.Country())

			if err != nil {
				return nil, errors.Join(err, tx.Rollback())
			}
		}

		province = province.ChangeCountry(province.Country().WithId(countryId))
	}

	id, err := repository.save(tx, province)
	if err != nil {
		return nil, errors.Join(err, tx.Rollback())
	}

	return province.WithId(id), tx.Commit()
}
