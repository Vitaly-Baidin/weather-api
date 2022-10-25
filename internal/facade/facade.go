package facade

import (
	"context"
	"github.com/Vitaly-Baidin/weather-api/entity"
)

type CityFacade interface {
	GetAll(ctx context.Context) ([]entity.CityResponse, error)
	GetSummary(ctx context.Context, country, name string) (entity.CityResponse, error)
}

type TemperatureFacade interface {
	GetWeatherDetail(ctx context.Context, country, name string, timestamp int) (entity.TemperatureResponse, error)
}
