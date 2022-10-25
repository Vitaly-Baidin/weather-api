// Package v1 of weather API
//
// Documentation for weather API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package v1

import "github.com/Vitaly-Baidin/weather-api/entity"

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// A list of city
// swagger:response cities
type citiesResponseWrapper struct {
	// All current products
	// in: body
	Body entity.Cities
}

// city
// swagger:response cityResponse
type cityResponseWrapper struct {
	// Newly created product
	// in: body
	Body entity.CityResponse
}

// temperature
// swagger:response temperatureResponse
type temperatureResponseWrapper struct {
	// Newly created product
	// in: body
	Body entity.TemperatureResponse
}
