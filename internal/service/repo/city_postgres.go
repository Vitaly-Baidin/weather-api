package repo

import (
	"context"
	"fmt"
	"github.com/Vitaly-Baidin/weather-api/entity"
	"github.com/Vitaly-Baidin/weather-api/internal/util"
	"github.com/Vitaly-Baidin/weather-api/pkg/postgres"
)

const (
	insertCityQuery = `INSERT INTO city (name, fmt_name, country, latitude, longitude) VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING `

	findCityByIDQuery            = `SELECT id, name, country, latitude, longitude FROM city WHERE id=$1`
	findCityByFullAddressQuery   = `SELECT id, name, country, latitude, longitude FROM city WHERE country=$1 AND fmt_name=$2`
	findCityIDByFullAddressQuery = `SELECT id FROM city WHERE country=$1 AND fmt_name=$2`
	findIDByCoordQuery           = `SELECT id FROM city WHERE latitude=$1 AND longitude=$2`
	findAllCoordQuery            = `SELECT latitude, longitude FROM city`

	ifExistsCityByCoordQuery = `SELECT exists(select 1 from city where latitude=$1 AND longitude=$2)`
)

// City -.
type City struct {
	*postgres.Postgres
}

// NewCity - .
func NewCity(pg *postgres.Postgres) *City {
	return &City{Postgres: pg}
}

func (r *City) Store(ctx context.Context, city entity.City) error {
	fmtName := util.FormatCityName(city.Name)
	fmtCountry := util.FormatCityName(city.Country)
	_, err := r.Pool.Exec(ctx, insertCityQuery, city.Name, fmtName, fmtCountry, city.Latitude, city.Longitude)
	if err != nil {
		return fmt.Errorf("repo.City - Store - Pool.Exec: %w", err)
	}

	return nil
}

func (r *City) FindByID(ctx context.Context, cityID uint) (entity.City, error) {
	row := r.Pool.QueryRow(ctx, findCityByIDQuery, cityID)
	result := entity.City{}

	err := row.Scan(&result.ID, &result.Name, &result.Country, &result.Latitude, &result.Longitude)
	if err != nil {
		return entity.City{}, fmt.Errorf("repo.City - FindByID - row.Scan: %w", err)
	}

	return result, nil
}

func (r *City) FindByFullAddress(ctx context.Context, country, name string) (entity.City, error) {
	row := r.Pool.QueryRow(ctx, findCityByFullAddressQuery, country, name)
	result := entity.City{}

	err := row.Scan(&result.ID, &result.Name, &result.Country, &result.Latitude, &result.Longitude)
	if err != nil {
		return entity.City{}, fmt.Errorf("repo.City - FindByFullAddress - row.Scan: %w", err)
	}

	return result, nil
}

func (r *City) FindIDByFullAddress(ctx context.Context, country, name string) (uint, error) {
	row := r.Pool.QueryRow(ctx, findCityIDByFullAddressQuery, country, name)
	var result uint

	err := row.Scan(&result)
	if err != nil {
		return 0, fmt.Errorf("repo.City - FindByFullAddress - row.Scan: %w", err)
	}

	return result, nil
}

func (r *City) FindIDByCoord(ctx context.Context, lat, lon float64) (uint, error) {
	row := r.Pool.QueryRow(ctx, findIDByCoordQuery, lat, lon)
	var result uint

	err := row.Scan(&result)
	if err != nil {
		return 0, fmt.Errorf("repo.City - FindIDByCoord - row.Scan: %w", err)
	}

	return result, nil
}

func (r *City) FindAllCoord(ctx context.Context) ([][2]float64, error) {
	rows, err := r.Pool.Query(ctx, findAllCoordQuery)
	if err != nil {
		return nil, fmt.Errorf("repo.City - FindAllCoord - Pool.Query: %w", err)
	}

	result := make([][2]float64, 0, 20)

	for rows.Next() {
		coord := [2]float64{}
		err = rows.Scan(&coord[0], &coord[1])
		if err != nil {
			return nil, fmt.Errorf("repo.City - FindAllCoord - rows.Scan: %w", err)
		}

		result = append(result, coord)
	}

	return result, nil
}

func (r *City) IfExistsByCoord(ctx context.Context, lat, lon float64) (bool, error) {
	row := r.Pool.QueryRow(ctx, ifExistsCityByCoordQuery, lat, lon)
	var result bool

	err := row.Scan(&result)
	if err != nil {
		return false, fmt.Errorf("repo.City - IfExistsByCoord - row.Scan: %w", err)
	}

	return result, nil
}
