package service

import (
	"context"
	"github.com/Vitaly-Baidin/weather-api/entity"
)

type (
	// City -.
	City interface {
		GetByID(ctx context.Context, cityID uint) (entity.City, error)
		GetByFullAddress(ctx context.Context, country, name string) (entity.City, error)
		GetIDByFullAddress(ctx context.Context, country, name string) (uint, error)
		SaveFromAPI(ctx context.Context, country, state, name string) error
	}

	// CityRepo -.
	CityRepo interface {
		Store(ctx context.Context, city entity.City) error
		FindByID(ctx context.Context, cityID uint) (entity.City, error)
		FindByFullAddress(ctx context.Context, country, name string) (entity.City, error)
		FindIDByFullAddress(ctx context.Context, country, name string) (uint, error)
		IfExistsByCoord(ctx context.Context, lat, lon float64) (bool, error)
	}

	// CityWebAPI -.
	CityWebAPI interface {
		FindByFullAddress(ctx context.Context, country, state, name string) (entity.City, error)
	}
)

type (
	// Temperature -.
	Temperature interface {
		GetAllUniqueCityID(ctx context.Context) ([]uint, error)
		GetActualMidTempByCityID(ctx context.Context, cityID uint) (float64, error)
		GetAllByCityID(ctx context.Context, cityID uint) ([]entity.Temperature, error)
		GetByCityIDAndTimestamp(ctx context.Context, cityID uint, timestamp int) (entity.TemperatureResponse, error)
		SaveFromAPI(ctx context.Context, lat, lon float64) error
	}

	// TemperatureRepo -.
	TemperatureRepo interface {
		FindAllUniqueCityID(ctx context.Context) ([]uint, error)
		FindActualMidTemp(ctx context.Context, cityID uint) (float64, error)
		FindAllByCityID(ctx context.Context, cityID uint) ([]entity.Temperature, error)
		FindByCityIDAndTimestamp(ctx context.Context, cityID uint, timestamp int) (entity.Temperature, error)
		Store(ctx context.Context, temperature entity.Temperature) error
	}

	// TemperatureWebAPI -.
	TemperatureWebAPI interface {
		FindByCoord(ctx context.Context, lat, lon float64) ([]entity.Temperature, error)
	}
)
