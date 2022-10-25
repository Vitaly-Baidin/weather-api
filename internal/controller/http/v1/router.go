package v1

import (
	"github.com/Vitaly-Baidin/weather-api/internal/facade"
	"github.com/Vitaly-Baidin/weather-api/pkg/logger"
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

	return sm
}
