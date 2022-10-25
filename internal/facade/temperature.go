package facade

import (
	"context"
	"fmt"
	"github.com/Vitaly-Baidin/weather-api/entity"
	"github.com/Vitaly-Baidin/weather-api/internal/service"
)

type Temperature struct {
	cityService service.City
	tempService service.Temperature
}

func NewTemperature(cityService service.City, tempService service.Temperature) *Temperature {
	return &Temperature{cityService: cityService, tempService: tempService}
}

func (f *Temperature) GetWeatherDetail(ctx context.Context, country, name string, timestamp int) (entity.TemperatureResponse, error) {
	id, err := f.cityService.GetIDByFullAddress(ctx, country, name)
	if err != nil {
		return entity.TemperatureResponse{}, fmt.Errorf("facade.Temperature - GetWeatherDetail - GetIDByFullAddress: %w", err)
	}

	temp, err := f.tempService.GetByCityIDAndTimestamp(ctx, id, timestamp)
	if err != nil {
		return entity.TemperatureResponse{}, fmt.Errorf("facade.Temperature - GetWeatherDetail - GetByCityIDAndTimestamp: %w", err)
	}

	return temp, nil
}
