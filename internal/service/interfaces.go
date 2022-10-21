package service

import (
	"context"
	"github.com/Vitaly-Baidin/weather-api/entity"
)

type (
	// City -.
	City interface {
		SetCity(ctx context.Context, name string) (entity.City, error)
		FindCityByName(ctx context.Context, name string) ([]entity.City, error)
		FindAllCities(ctx context.Context) ([]entity.City, error)
	}

	// CityRepo -.
	CityRepo interface {
		SaveCity(ctx context.Context, entity entity.City) error
		GetCityByName(ctx context.Context, name string) ([]entity.City, error)
		GetAllCities(ctx context.Context) ([]entity.City, error)
		IfExists(ctx context.Context, city entity.City) (bool, error)
	}

	// CityWebAPI -.
	CityWebAPI interface {
		GetInformation(name string) ([]entity.City, error)
	}
)
