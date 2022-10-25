package entity

// Cities defines the structure for an weather API
// swagger:model
type Cities []City

// City defines the structure for an weather API
// swagger:model
type City struct {
	// the id for the City
	//
	// required: false
	// min: 1
	ID uint `json:"-"`
	// the name for this City
	//
	// required: true
	// max length: 200
	Name string `json:"name"`
	// the name country for this City
	//
	// required: true
	// max length: 200
	Country string `json:"country"`
	// the latitude for this City
	//
	// required: true
	Latitude float64 `json:"lat"`
	// the longitude for this City
	//
	// required: true
	Longitude float64 `json:"lon"`
}

// CityResponse defines the structure for an weather API
// swagger:model
type CityResponse struct {
	// the name for this CityResponse
	Name string `json:"name"`
	// the name country for this CityResponse
	Country string `json:"country"`
	// actual temp for this CityResponse
	Temperature float64 `json:"temperature,omitempty"`
	// list temp for this CityResponse
	Weather []Temperature `json:"weather,omitempty"`
}
