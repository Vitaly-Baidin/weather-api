package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/Vitaly-Baidin/weather-api/entity"
)

// CityService -.
type CityService struct {
	repo   CityRepo
	webAPI CityWebAPI
}

// NewCity -.
func NewCity(r CityRepo, api CityWebAPI) *CityService {
	return &CityService{repo: r, webAPI: api}
}

// SetCity -.
func (s *CityService) SetCity(ctx context.Context, name string) (entity.City, error) {
	cities, err := s.webAPI.GetInformation(name)
	if err != nil {
		return entity.City{}, fmt.Errorf("CityService - SetCity - webAPI.GetInformation: %w", err)
	}

	flag := -1

	for i, city := range cities {
		ifExists, err := s.repo.IfExists(ctx, city)
		if err != nil {
			return entity.City{}, fmt.Errorf("CityService - SetCity - repo.IfExists: %w", err)
		}
		if !ifExists {
			flag = i
			break
		}
	}

	if flag == -1 {
		return entity.City{}, errors.New("city already exists")
	}

	err = s.repo.SaveCity(ctx, cities[flag])
	if err != nil {
		return entity.City{}, fmt.Errorf("CityService - SetCity - repo.SaveCity: %w", err)
	}

	return cities[flag], nil
}

// FindCityByName -.
func (s *CityService) FindCityByName(ctx context.Context, name string) ([]entity.City, error) {
	cities, err := s.repo.GetCityByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("CityService - FindCityByName - repo.GetCityByName: %w", err)
	}
	return cities, nil
}

// FindAllCities -.
func (s *CityService) FindAllCities(ctx context.Context) ([]entity.City, error) {
	cities, err := s.repo.GetAllCities(ctx)
	if err != nil {
		return nil, fmt.Errorf("CityService - FindAllCities - repo.GetAllCities: %w", err)
	}
	return cities, nil
}
