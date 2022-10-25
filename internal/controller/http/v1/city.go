package v1

import (
	"errors"
	"github.com/Vitaly-Baidin/weather-api/entity"
	"github.com/Vitaly-Baidin/weather-api/internal/facade"
	"github.com/Vitaly-Baidin/weather-api/internal/util"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"net/http"
)

type City struct {
	city facade.CityFacade
}

func NewCity(city facade.CityFacade) *City {
	return &City{city: city}
}

func (h *City) ListAll(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	cities, err := h.city.GetAll(r.Context())
	if errors.Is(err, pgx.ErrNoRows) {
		rw.WriteHeader(http.StatusNotFound)
		entity.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		entity.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = entity.ToJSON(cities, rw)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		entity.ToJSON(&GenericError{Message: err.Error()}, rw)
	}
}

func (h *City) SummarySingle(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	countryName := util.FormatCityName(vars["country"])
	cityName := util.FormatCityName(vars["name"])

	city, err := h.city.GetSummary(r.Context(), countryName, cityName)
	if errors.Is(err, pgx.ErrNoRows) {
		rw.WriteHeader(http.StatusNotFound)
		entity.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	} else if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		entity.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}
	err = entity.ToJSON(city, rw)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		entity.ToJSON(&GenericError{Message: err.Error()}, rw)
	}
}
