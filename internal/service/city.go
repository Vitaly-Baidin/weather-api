package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/Vitaly-Baidin/weather-api/entity"
)

var errCityExists = errors.New("city already exists")

// CityService -.
type CityService struct {
	repo   CityRepo
	webAPI CityWebAPI
}

// NewCity -.
func NewCity(r CityRepo, api CityWebAPI) *CityService {
	return &CityService{repo: r, webAPI: api}
}

func (s *CityService) GetByID(ctx context.Context, cityID uint) (entity.City, error) {
	city, err := s.repo.FindByID(ctx, cityID)
	if err != nil {
		return entity.City{}, fmt.Errorf("service.CityService - GetByID - FindByID: %w", err)
	}

	return city, nil
}

func (s *CityService) GetByFullAddress(ctx context.Context, country, name string) (entity.City, error) {
	city, err := s.repo.FindByFullAddress(ctx, country, name)
	if err != nil {
		return entity.City{}, fmt.Errorf("service.CityService - GetByFullAddress - FindByFullAddress: %w", err)
	}

	return city, nil
}

func (s *CityService) GetIDByFullAddress(ctx context.Context, country, name string) (uint, error) {
	id, err := s.repo.FindIDByFullAddress(ctx, country, name)
	if err != nil {
		return 0, fmt.Errorf("service.CityService - GetIDByFullAddress - repo.FindIDByFullAddress: %w", err)
	}

	return id, nil
}

func (s *CityService) SaveFromAPI(ctx context.Context, country, state, name string) error {
	city, err := s.webAPI.FindByFullAddress(ctx, country, state, name)
	if err != nil {
		return fmt.Errorf("CityService - SaveFromAPI - webAPI.FindByFullAddress: %w", err)
	}

	ifExists, err := s.repo.IfExistsByCoord(ctx, city.Longitude, city.Latitude)
	if err != nil {
		return fmt.Errorf("CityService - SaveFromAPI - repo.IfExistsByCoord: %w", err)
	}

	if ifExists {
		return errCityExists
	}

	err = s.repo.Store(ctx, city)
	if err != nil {
		return fmt.Errorf("CityService - SaveFromAPI - repo.Store: %w", err)
	}

	return nil
}
