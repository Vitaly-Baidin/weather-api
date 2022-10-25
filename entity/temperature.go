package entity

import (
	"encoding/json"
)

// Temperatures defines the structure for an weather API
// swagger:model
type Temperatures struct {
	// list Temperature
	Data []json.RawMessage `json:"list"`
}

// Temperature defines the structure for an weather API
// swagger:model
type Temperature struct {
	// timestamp for the temperature
	//
	// required: true
	Timestamp int64 `json:"dt"`
	// temperature for the temperature
	//
	// required: true
	Temperature float64 `njson:"main.temp"`
	// cityID for the temperature
	//
	// required: true
	CityID uint `json:"-"`
	// data for the temperature
	//
	// required: true
	Data []byte `json:"-"`
}

// TemperatureResponse defines the structure for an weather API
// swagger:model
type TemperatureResponse struct {
	// the timestamp for this TemperatureResponse
	Timestamp int64 `json:"date"`
	// the temperature for this TemperatureResponse
	Temperature float64 `json:"temperature"`
	// the data for this TemperatureResponse
	Data json.RawMessage `json:"data"`
}
