package repo

import (
	"context"
	"fmt"
	"github.com/Vitaly-Baidin/weather-api/entity"
	"github.com/Vitaly-Baidin/weather-api/pkg/postgres"
	"time"
)

const (
	insertTemperatureQuery = `INSERT INTO temperature (timestamp, temp, city_id, data) 
							  VALUES ($1, $2, $3, $4) 
							  ON CONFLICT (timestamp, city_id) DO UPDATE 
							  SET temp=excluded.temp, data=excluded.data`

	findAllUniqueCityIDQuery      = `SELECT DISTINCT city_id FROM temperature`
	findActualMidTempQuery        = `SELECT AVG(temp) FROM temperature WHERE city_id=$1 AND timestamp >= $2`
	findAllByCityIDQuery          = `SELECT timestamp, temp, city_id, data FROM temperature WHERE city_id=$1`
	findByCityIDAndTimestampQuery = `SELECT timestamp, temp, city_id, data FROM temperature WHERE city_id=$1 AND timestamp=$2`
)

type Temperature struct {
	*postgres.Postgres
}

func NewTemperature(db *postgres.Postgres) *Temperature {
	return &Temperature{Postgres: db}
}

func (r *Temperature) FindAllUniqueCityID(ctx context.Context) ([]uint, error) {
	rows, err := r.Pool.Query(ctx, findAllUniqueCityIDQuery)
	if err != nil {
		return nil, fmt.Errorf("Temperature - FindAllUniqueCityID - Pool.Query: %w", err)
	}

	result := make([]uint, 0, 128)

	for rows.Next() {
		var cityID uint
		err = rows.Scan(&cityID)
		if err != nil {
			return nil, fmt.Errorf("Temperature - FindAllUniqueCityID - rows.Scan: %w", err)
		}
		result = append(result, cityID)
	}

	return result, nil
}

func (r *Temperature) FindActualMidTemp(ctx context.Context, cityID uint) (float64, error) {
	var result float64

	actualTime := time.Now().Unix()

	row := r.Pool.QueryRow(ctx, findActualMidTempQuery, cityID, actualTime)
	err := row.Scan(&result)
	if err != nil {
		return 0, fmt.Errorf("Temperature - FindActualMidTemp - Scan: %w", err)
	}

	return result, nil
}

func (r *Temperature) FindAllByCityID(ctx context.Context, cityID uint) ([]entity.Temperature, error) {
	rows, err := r.Pool.Query(ctx, findAllByCityIDQuery, cityID)
	if err != nil {
		return nil, fmt.Errorf("Temperature - FindAllByCityID - Pool.Query: %w", err)
	}

	result := make([]entity.Temperature, 0, 40)

	for rows.Next() {
		temp := entity.Temperature{}
		err := rows.Scan(&temp.Timestamp, &temp.Temperature, &temp.CityID, &temp.Data)
		if err != nil {
			return nil, fmt.Errorf("Temperature - FindAllByCityID - rows.Scan: %w", err)
		}

		result = append(result, temp)
	}

	return result, nil
}

func (r *Temperature) FindByCityIDAndTimestamp(ctx context.Context, cityID uint, timestamp int) (entity.Temperature, error) {
	temp := entity.Temperature{}

	row := r.Pool.QueryRow(ctx, findByCityIDAndTimestampQuery, cityID, timestamp)
	err := row.Scan(&temp.Timestamp, &temp.Temperature, &temp.CityID, &temp.Data)
	if err != nil {
		return entity.Temperature{}, fmt.Errorf("Temperature - FindByCityIDAndTimestamp - rows.Scan: %w", err)
	}

	return temp, nil
}

func (r *Temperature) Store(ctx context.Context, temperature entity.Temperature) error {
	_, err := r.Pool.Exec(ctx, insertTemperatureQuery, temperature.Timestamp, temperature.Temperature, temperature.CityID, temperature.Data)
	if err != nil {
		return fmt.Errorf("temperature - Store - Pool.Exec: %w", err)
	}

	return nil
}
