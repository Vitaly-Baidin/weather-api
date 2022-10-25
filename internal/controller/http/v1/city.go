package v1

import (
	"errors"
	"github.com/Vitaly-Baidin/weather-api/entity"
	"github.com/Vitaly-Baidin/weather-api/internal/facade"
	"github.com/Vitaly-Baidin/weather-api/internal/util"
	"github.com/Vitaly-Baidin/weather-api/pkg/logger"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"net/http"
)

type City struct {
	city facade.CityFacade
	log  logger.Logger
}

func NewCity(city facade.CityFacade, log logger.Logger) *City {
	return &City{city: city, log: log}
}

func (h *City) ListAll(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	cities, err := h.city.GetAll(r.Context())
	if errors.Is(err, pgx.ErrNoRows) {
		h.log.Error(err, "http - v1 - ListAll")
		rw.WriteHeader(http.StatusNotFound)
		entity.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}
	if err != nil {
		h.log.Error(err, "http - v1 - ListAll")
		rw.WriteHeader(http.StatusInternalServerError)
		entity.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = entity.ToJSON(cities, rw)
	if err != nil {
		h.log.Error(err, "http - v1 - ListAll")
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
		h.log.Error(err, "http - v1 - SummarySingle")
		rw.WriteHeader(http.StatusNotFound)
		entity.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	} else if err != nil {
		h.log.Error(err, "http - v1 - SummarySingle")
		rw.WriteHeader(http.StatusInternalServerError)
		entity.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}
	err = entity.ToJSON(city, rw)
	if err != nil {
		h.log.Error(err, "http - v1 - SummarySingle")
		rw.WriteHeader(http.StatusInternalServerError)
		entity.ToJSON(&GenericError{Message: err.Error()}, rw)
	}
}
