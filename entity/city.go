package entity

type Cities []City

// City defines the structure for an weather API
// swagger:model
type City struct {
	// the id for the city
	//
	// required: false
	// min: 1
	ID uint `json:"-"`
	// the name for this city
	//
	// required: true
	// max length: 200
	Name string `json:"name"`
	// the name country for this city
	//
	// required: true
	// max length: 200
	Country string `json:"country"`
	// the latitude for this city
	//
	// required: true
	Latitude float64 `json:"lat"`
	// the longitude for this city
	//
	// required: true
	Longitude float64 `json:"lon"`
}

// CityResponse defines the structure for an weather API
// swagger:model
type CityResponse struct {
	// the name for this city
	Name string `json:"name"`
	// the name country for this city
	Country string `json:"country"`
	// actual temp for this city
	Temperature float64 `json:"temperature,omitempty"`
	// list temp for this city
	Weather []Temperature `json:"weather,omitempty"`
}
