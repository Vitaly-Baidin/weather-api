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

// CityInformationWebAPI -.
type CityInformationWebAPI struct {
}

// New -.
func New() *CityInformationWebAPI {
	return &CityInformationWebAPI{}
}

func (wa *CityInformationWebAPI) GetInformation(name string) (entity.City, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&limit=1&appid=%s", name, apikey)
	response, err := http.Get(url)
	if err != nil {
		return entity.City{}, fmt.Errorf("CityInformationWebAPI = GetInformation - http.Get: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return entity.City{}, fmt.Errorf("CityInformationWebAPI = GetInformation - io.ReadAll: %w", err)
	}

	var cities entity.Cities

	err = json.Unmarshal(body, &cities)
	if err != nil {
		return entity.City{}, fmt.Errorf("CityInformationWebAPI = GetInformation - json.Unmarshal: %w", err)
	}

	if len(cities) == 0 {
		return entity.City{}, fmt.Errorf("city not found")
	}

	return cities[0], nil
}
