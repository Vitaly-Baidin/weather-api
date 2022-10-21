package repo

import (
	"context"
	"fmt"
	"github.com/Vitaly-Baidin/weather-api/entity"
	"github.com/Vitaly-Baidin/weather-api/pkg/postgres"
)

const (
	defaultEntityCap = 64

	insertCityQuery     = `INSERT INTO city (name, country, latitude, longitude) VALUES ($1, $2, $3, $4)`
	findCityByNameQuery = `SELECT name, country, latitude, longitude FROM city WHERE name='$1'`
	findAllCitiesQuery  = `SELECT name, country, latitude, longitude FROM city`
	cityIfExistsQuery   = `SELECT exists(select 1 from city where latitude=$1 and longitude=$2)`
)

// CityRepo -.
type CityRepo struct {
	*postgres.Postgres
}

// NewCity -.
func NewCity(pg *postgres.Postgres) *CityRepo {
	return &CityRepo{Postgres: pg}
}

// SaveCity -.
func (r *CityRepo) SaveCity(ctx context.Context, entity entity.City) error {
	_, err := r.Pool.Exec(ctx, insertCityQuery, entity.Name, entity.Country, entity.Latitude, entity.Longitude)
	if err != nil {
		return fmt.Errorf("CityRepo - SaveCity - r.Pool.Exec: %w", err)
	}

	return nil
}

// GetCityByName -.
func (r *CityRepo) GetCityByName(ctx context.Context, name string) ([]entity.City, error) {
	rows, err := r.Pool.Query(ctx, findCityByNameQuery, name)
	if err != nil {
		return nil, fmt.Errorf("CityRepo - FindCityByName - r.Pool.Query: %w", err)
	}

	result := make(entity.Cities, 0, defaultEntityCap)

	for rows.Next() {
		city := entity.City{}

		err = rows.Scan(&city.Name, &city.Country, &city.Latitude, &city.Longitude)
		if err != nil {
			return nil, fmt.Errorf("CityRepo - FindCityByName - rows.Scan: %w", err)
		}

		result = append(result)
	}

	return result, nil
}

// GetAllCities -.
func (r *CityRepo) GetAllCities(ctx context.Context) ([]entity.City, error) {
	rows, err := r.Pool.Query(ctx, findAllCitiesQuery)
	if err != nil {
		return nil, fmt.Errorf("CityRepo - FindAllCities - r.Pool.Query: %w", err)
	}

	result := make(entity.Cities, 0, defaultEntityCap)

	for rows.Next() {
		city := entity.City{}

		err = rows.Scan(&city.Name, &city.Country, &city.Latitude, &city.Longitude)
		if err != nil {
			return nil, fmt.Errorf("CityRepo - FindAllCities - rows.Scan: %w", err)
		}

		result = append(result, city)
	}

	return result, nil
}

// IfExists -.
func (r *CityRepo) IfExists(ctx context.Context, city entity.City) (bool, error) {
	var result bool
	err := r.Pool.QueryRow(ctx, cityIfExistsQuery, city.Latitude, city.Longitude).Scan(&result)
	if err != nil {
		return false, fmt.Errorf("CityRepo - IfExists - r.Pool.QueryRow: %w", err)
	}
	return result, nil
}
