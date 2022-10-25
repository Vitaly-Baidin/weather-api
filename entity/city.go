package entity

type Cities []City

type City struct {
	ID        uint    `json:"-"`
	Name      string  `json:"name"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

type CityResponse struct {
	Name        string        `json:"name"`
	Country     string        `json:"country"`
	Temperature float64       `json:"temperature,omitempty"`
	Weather     []Temperature `json:"weather,omitempty"`
	Links       []Link        `json:"links,omitempty"`
}
