package webapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Vitaly-Baidin/weather-api/entity"
	"github.com/m7shapan/njson"
	"io"
	"net/http"
	"net/url"
	"time"
)

// TemperatureWebAPI -.
type TemperatureWebAPI struct {
}

// NewTemperature -.
func NewTemperature() *TemperatureWebAPI {
	return &TemperatureWebAPI{}
}

// FindByCoord -.
func (wa *TemperatureWebAPI) FindByCoord(ctx context.Context, lat float64, lon float64) ([]entity.Temperature, error) {
	queries := make(url.Values)

	queries.Add("lat", fmt.Sprintf("%f", lat))
	queries.Add("lon", fmt.Sprintf("%f", lon))
	queries.Add("appid", apikey)
	queries.Add("units", "metric")

	u := url.URL{
		Scheme:   "http",
		Host:     "api.openweathermap.org",
		Path:     "data/2.5/forecast",
		RawQuery: queries.Encode(),
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	r, err := client.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("TemperatureWebAPI - FindByCoord - http.Get: %w", err)
	}
	defer r.Body.Close()

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("TemperatureWebAPI - FindByCoord - io.ReadAll: %w", err)
	}

	response := entity.Temperatures{}

	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, fmt.Errorf("TemperatureWebAPI - FindByCoord - json.Unmarshal: %w", err)
	}

	var result []entity.Temperature

	for _, msg := range response.Data {
		var temp entity.Temperature
		tempJSON, err := msg.MarshalJSON()
		if err != nil {
			return nil, fmt.Errorf("TemperatureWebAPI - FindByCoord - json.Unmarshal: %w", err)
		}

		err = njson.Unmarshal(tempJSON, &temp)
		if err != nil {
			return nil, fmt.Errorf("TemperatureWebAPI - FindByCoord - json.Unmarshal: %w", err)
		}

		temp.Data = append(temp.Data, tempJSON...)
		result = append(result, temp)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("city not found")
	}

	return result, nil
}
