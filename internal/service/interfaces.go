package service

import "github.com/Vitaly-Baidin/weather-api/entity"

// CityInformationWebAPI -.
type CityInformationWebAPI interface {
	GetInformation(name string) (entity.City, error)
}
