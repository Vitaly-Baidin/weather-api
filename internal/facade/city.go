package facade

import (
	"context"
	"fmt"
	"github.com/Vitaly-Baidin/weather-api/entity"
	"github.com/Vitaly-Baidin/weather-api/internal/service"
	"strings"
)

type City struct {
	cityService service.City
	tempService service.Temperature
}

func NewCity(cityService service.City, tempService service.Temperature) *City {
	return &City{cityService: cityService, tempService: tempService}
}

func (f *City) GetAll(ctx context.Context) ([]entity.CityResponse, error) {
	citiesID, err := f.tempService.GetAllUniqueCityID(ctx)
	if err != nil {
		return nil, fmt.Errorf("facade.City - GetAllCities - GetAllUniqueCityID: %w", err)
	}

	result := make([]entity.CityResponse, 0, 20)

	for _, id := range citiesID {
		city, err := f.cityService.GetByID(ctx, id)
		if err != nil {
			return nil, fmt.Errorf("facade.City - GetAllCities - GetByID: %w", err)
		}

		response := entity.CityResponse{
			Name:    city.Name,
			Country: city.Country,
		}
		result = append(result, response)
	}

	return result, nil
}

func (f *City) GetSummary(ctx context.Context, country, name string) (entity.CityResponse, error) {
	city, err := f.cityService.GetByFullAddress(ctx, country, name)
	if err != nil {
		return entity.CityResponse{}, fmt.Errorf("facade.City - GetSummary - FindCityByFullAddress: %w", err)
	}

	midTemp, err := f.tempService.GetActualMidTempByCityID(ctx, city.ID)
	if err != nil {
		return entity.CityResponse{}, fmt.Errorf("facade.City - GetSummary - GetActualMidTempByCityID: %w", err)
	}

	temperatures, err := f.tempService.GetAllByCityID(ctx, city.ID)
	if err != nil {
		return entity.CityResponse{}, fmt.Errorf("facade.City - GetSummary - GetAllByCityID: %w", err)
	}

	response := createCityResponse(city, midTemp, temperatures)

	return response, nil

}

func createCityResponse(city entity.City, temperature float64, weather []entity.Temperature) entity.CityResponse {
	links := []entity.Link{
		{
			Href: fmt.Sprintf("city/%s/%s/summary", strings.ToLower(city.Country), strings.ToLower(city.Name)),
			Type: "GET",
		},
		{
			Href: fmt.Sprintf("city/%s/%s/detail", strings.ToLower(city.Country), strings.ToLower(city.Name)),
			Type: "GET",
		},
	}
	cr := entity.CityResponse{
		Name:        city.Name,
		Country:     city.Country,
		Temperature: temperature,
		Weather:     weather,
		Links:       links,
	}

	return cr
}
