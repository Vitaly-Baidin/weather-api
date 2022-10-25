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

import (
	"github.com/Vitaly-Baidin/weather-api/internal/facade"
	"github.com/Vitaly-Baidin/weather-api/pkg/logger"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(cityFacade facade.CityFacade, tempFacade facade.TemperatureFacade, l logger.Logger) *mux.Router {
	sm := mux.NewRouter()

	cf := NewCity(cityFacade, l)
	tf := NewTemperature(tempFacade, l)

	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/city", cf.ListAll)
	getR.HandleFunc("/city/{country}/{name}/summary", cf.SummarySingle)
	getR.HandleFunc("/city/{country}/{name}/weather/{timestamp}", tf.DetailSingle)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getR.Handle("/docs", sh)
	getR.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	return sm
}
