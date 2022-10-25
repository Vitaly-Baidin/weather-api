package service

import (
	"context"
	"fmt"
	"github.com/Vitaly-Baidin/weather-api/entity"
)

type TemperatureService struct {
	repo   TemperatureRepo
	webAPI TemperatureWebAPI
}

func NewTemperature(repo TemperatureRepo, webAPI TemperatureWebAPI) *TemperatureService {
	return &TemperatureService{repo: repo, webAPI: webAPI}
}

func (s *TemperatureService) GetAllUniqueCityID(ctx context.Context) ([]uint, error) {
	citiesID, err := s.repo.FindAllUniqueCityID(ctx)
	if err != nil {
		return nil, fmt.Errorf("TemperatureService - GetAllUniqueCityID - FindAllUniqueCityID: %w", err)
	}

	return citiesID, nil
}

func (s *TemperatureService) GetActualMidTempByCityID(ctx context.Context, cityID uint) (float64, error) {
	actualMidTemp, err := s.repo.FindActualMidTemp(ctx, cityID)
	if err != nil {
		return 0, fmt.Errorf("TemperatureService - GetAllUniqueCityID - FindAllUniqueCityID: %w", err)
	}

	return actualMidTemp, nil
}

func (s *TemperatureService) GetAllByCityID(ctx context.Context, cityID uint) ([]entity.Temperature, error) {
	temperatures, err := s.repo.FindAllByCityID(ctx, cityID)
	if err != nil {
		return nil, fmt.Errorf("TemperatureService - GetAllUniqueCityID - FindAllUniqueCityID: %w", err)
	}

	return temperatures, nil
}

func (s *TemperatureService) GetByCityIDAndTimestamp(ctx context.Context, cityID uint, timestamp int) (entity.TemperatureResponse, error) {
	temperature, err := s.repo.FindByCityIDAndTimestamp(ctx, cityID, timestamp)
	if err != nil {
		return entity.TemperatureResponse{}, fmt.Errorf("TemperatureService - GetAllUniqueCityID - FindAllUniqueCityID: %w", err)
	}

	result := entity.TemperatureResponse{
		Temperature: temperature.Temperature,
		Timestamp:   temperature.Timestamp,
		Data:        temperature.Data,
	}

	return result, nil
}

func (s *TemperatureService) SaveFromAPI(ctx context.Context, lat, lon float64) error {
	temps, err := s.webAPI.FindByCoord(ctx, lat, lon)
	if err != nil {
		return fmt.Errorf("TemperatureService - SaveFromAPI - webAPI.FindByCoord: %w", err)
	}

	for _, temp := range temps {
		err = s.repo.Store(ctx, temp)
		if err != nil {
			return fmt.Errorf("TemperatureService - SaveFromAPI - repo.Store: %w", err)
		}
	}

	return nil
}
