package v1

import (
	"errors"
	"github.com/Vitaly-Baidin/weather-api/entity"
	"github.com/Vitaly-Baidin/weather-api/internal/facade"
	"github.com/Vitaly-Baidin/weather-api/pkg/logger"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"net/http"
	"strconv"
)

type Temperature struct {
	temp facade.TemperatureFacade
	log  logger.Logger
}

func NewTemperature(temp facade.TemperatureFacade, log logger.Logger) *Temperature {
	return &Temperature{temp: temp, log: log}
}

// swagger:route GET /city/{country}/{name}/weather/{timestamp} temperature DetailSingle
// Return temperature from the database
// responses:
//	200: TemperatureResponse
//	404: GenericError
//	500: GenericError

// DetailSingle handles GET requests
func (h *Temperature) DetailSingle(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)

	countryName := vars["country"]
	cityName := vars["name"]
	timestamp := vars["timestamp"]

	tsInt, err := strconv.Atoi(timestamp)
	if err != nil {
		h.log.Error(err, "http - v1 - DetailSingle")
		rw.WriteHeader(http.StatusBadRequest)
		entity.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	cities, err := h.temp.GetWeatherDetail(r.Context(), countryName, cityName, tsInt)
	if errors.Is(err, pgx.ErrNoRows) {
		h.log.Error(err, "http - v1 - DetailSingle")
		rw.WriteHeader(http.StatusNotFound)
		entity.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	} else if err != nil {
		h.log.Error(err, "http - v1 - DetailSingle")
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
