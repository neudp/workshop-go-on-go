package addressRepository

import (
	"database/sql"
	_ "embed"
	"errors"
	"goOnGo/internal/orders/model/domain/address"
)

var (
	//go:embed query/insert-country.sql
	insertCountryQuery string
	//go:embed query/select-country-id-by-name.sql
	selectCountryIdByNameQuery string
)

type CountryRepository struct {
	db *sql.DB
}

func NewCountryRepository(db *sql.DB) *CountryRepository {
	return &CountryRepository{db: db}
}

func (*CountryRepository) save(
	tx *sql.Tx,
	country *address.Country,
) (address.CountryId, error) {
	stmt, err := tx.Exec(insertCountryQuery, country.Name())

	if err != nil {
		return 0, err
	}

	id, err := stmt.LastInsertId()
	if err != nil {
		return 0, err
	}

	return address.CountryId(id), err
}

func (*CountryRepository) findIdByName(
	tx *sql.Tx,
	countryName string,
) (address.CountryId, error) {
	var id int64
	err := tx.QueryRow(selectCountryIdByNameQuery, countryName).Scan(&id)

	if err != nil {
		return 0, err
	}

	return address.CountryId(id), nil
}

func (repository *CountryRepository) Save(country *address.Country) (*address.Country, error) {
	tx, err := repository.db.Begin()
	if err != nil {
		return nil, err
	}

	id, err := repository.save(tx, country)
	if err != nil {
		return nil, errors.Join(err, tx.Rollback())
	}

	return country.WithId(id), tx.Commit()
}
