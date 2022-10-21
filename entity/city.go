package entity

type Cities []City

type City struct {
	Name      string  `json:"name"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

//type CityDTO struct {
//	ID        string  `json:"id"`
//	Name      string  `json:"name"`
//	Country   string  `json:"country"`
//	Latitude  float64 `json:"lat"`
//	Longitude float64 `json:"lon"`
//}
