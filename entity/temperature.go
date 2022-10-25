package entity

import (
	"encoding/json"
)

type Temperatures struct {
	Data []json.RawMessage `json:"list"`
}

type Temperature struct {
	Timestamp   int64   `json:"dt"`
	Temperature float64 `njson:"main.temp"`
	CityID      uint    `json:"-"`
	Data        []byte  `json:"-"`
}

type TemperatureResponse struct {
	Timestamp   int64   `json:"date"`
	Temperature float64 `json:"temperature"`
	Data        []byte  `json:"data"`
}
