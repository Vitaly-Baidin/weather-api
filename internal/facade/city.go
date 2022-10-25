package facade

import (
	"context"
	"fmt"
	"github.com/Vitaly-Baidin/weather-api/entity"
	"github.com/Vitaly-Baidin/weather-api/internal/service"
	"golang.org/x/sync/errgroup"
	"sort"
	"time"
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

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
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

func (f *City) UpdateActualTemp(ctx context.Context) error {
	allCoord, err := f.cityService.GetAllCoord(ctx)
	if err != nil {
		return fmt.Errorf("facade.City - UpdateActualTemp - GetAllCoord: %w", err)
	}

	g, ctx := errgroup.WithContext(ctx)

	for _, elem := range allCoord {
		cityID, err := f.cityService.GetIDByCoord(ctx, elem[0], elem[1])
		if err != nil {
			return fmt.Errorf("facade.City - UpdateActualTemp - GetIDByCoord: %w", err)
		}

		g.Go(func() error {
			time.Sleep(200 * time.Millisecond)
			err = f.tempService.SaveFromAPI(ctx, elem[0], elem[1], cityID)
			if err != nil {
				return err
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("facade.City - UpdateActualTemp - SaveFromAPI: %w", err)
	}

	return nil
}

func createCityResponse(city entity.City, temperature float64, weather []entity.Temperature) entity.CityResponse {
	cr := entity.CityResponse{
		Name:        city.Name,
		Country:     city.Country,
		Temperature: temperature,
		Weather:     weather,
	}

	return cr
}
