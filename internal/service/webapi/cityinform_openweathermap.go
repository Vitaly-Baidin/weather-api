package webapi

import (
	"encoding/json"
	"fmt"
	"github.com/Vitaly-Baidin/weather-api/entity"
	"io"
	"net/http"
)

var (
	apikey = "c4435be349344355dfeeadc4261d8e59"
)

// CityWebAPI -.
type CityWebAPI struct {
}

// NewCity -.
func NewCity() *CityWebAPI {
	return &CityWebAPI{}
}

func (wa *CityWebAPI) GetInformation(name string) ([]entity.City, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&appid=%s", name, apikey) // ?? &limit=1
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("CityWebAPI = GetInformation - http.Get: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("CityWebAPI = GetInformation - io.ReadAll: %w", err)
	}

	var cities entity.Cities

	err = json.Unmarshal(body, &cities)
	if err != nil {
		return nil, fmt.Errorf("CityWebAPI = GetInformation - json.Unmarshal: %w", err)
	}

	if len(cities) == 0 {
		return nil, fmt.Errorf("city not found")
	}

	return cities, nil
}
