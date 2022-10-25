package webapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Vitaly-Baidin/weather-api/entity"
	"io"
	"net/http"
	"net/url"
	"time"
)

// CityWebAPI -.
type CityWebAPI struct {
	apiKey string
}

// NewCity -.
func NewCity(apiKey string) *CityWebAPI {
	return &CityWebAPI{apiKey: apiKey}
}

func (wa *CityWebAPI) FindByFullAddress(ctx context.Context, country, state, name string) (entity.City, error) {
	queries := make(url.Values)

	queries.Add("q", createNameStr(country, state, name))
	queries.Add("limit", "1")
	queries.Add("appid", wa.apiKey)

	u := url.URL{
		Scheme:   "http",
		Host:     "api.openweathermap.org",
		Path:     "geo/1.0/direct",
		RawQuery: queries.Encode(),
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	response, err := client.Get(u.String())
	if err != nil {
		return entity.City{}, fmt.Errorf("CityWebAPI - GetInformationByName - http.Get: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return entity.City{}, fmt.Errorf("CityWebAPI - GetInformationByName - io.ReadAll: %w", err)
	}

	var cities entity.Cities

	err = json.Unmarshal(body, &cities)
	if err != nil {
		return entity.City{}, fmt.Errorf("CityWebAPI - GetInformationByName - json.Unmarshal: %w", err)
	}

	if len(cities) == 0 {
		return entity.City{}, fmt.Errorf("city not found")
	}

	return cities[0], nil
}

func createNameStr(country, state, name string) string {
	if state == "" && country == "" {
		return name
	} else if state == "" {
		return fmt.Sprintf("%s,%s", name, country)
	}

	return fmt.Sprintf("%s,%s,%s", name, state, country)
}
