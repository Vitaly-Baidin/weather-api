package app

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/Vitaly-Baidin/weather-api/config"
	v1 "github.com/Vitaly-Baidin/weather-api/internal/controller/http/v1"
	"github.com/Vitaly-Baidin/weather-api/internal/cronjob"
	"github.com/Vitaly-Baidin/weather-api/internal/facade"
	"github.com/Vitaly-Baidin/weather-api/internal/service"
	"github.com/Vitaly-Baidin/weather-api/internal/service/repo"
	"github.com/Vitaly-Baidin/weather-api/internal/service/webapi"
	"github.com/Vitaly-Baidin/weather-api/pkg/httpserver"
	"github.com/Vitaly-Baidin/weather-api/pkg/logger"
	"github.com/Vitaly-Baidin/weather-api/pkg/postgres"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	db, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	defer db.Close()
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}

	// Register Repositories
	cityRepo := repo.NewCity(db)
	tempRepo := repo.NewTemperature(db)

	// Register WebAPIs
	cityAPI := webapi.NewCity(cfg.Api.OpenweathermapKey)
	tempAPI := webapi.NewTemperature(cfg.Api.OpenweathermapKey)

	// Register Services
	cityService := service.NewCity(cityRepo, cityAPI)
	tempService := service.NewTemperature(tempRepo, tempAPI)

	// Register Facades
	cityFacade := facade.NewCity(cityService, tempService)
	tempFacade := facade.NewTemperature(cityService, tempService)

	// Examples cities
	initCities("city.csv", cityService)

	// Register cron
	cron := cronjob.NewCron(l)
	cron.StartCron()

	cronjob.RegisterJob(cron, func() {
		ctx := context.Background()
		err := cityFacade.UpdateActualTemp(ctx)
		if err != nil {
			l.Error(err, "app - RegisterJob - UpdateActualTemp")
			return
		}
		l.Info("UpdateActualTemp")
	})

	defer cron.StopCron()

	// Register Server
	rv1 := v1.NewRouter(cityFacade, tempFacade, l)

	httpServer := httpserver.New(rv1, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}

func initCities(filename string, city service.City) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("Unable to read input file "+filename, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filename, err)
	}

	for _, elem := range records {
		time.Sleep(200 * time.Millisecond)
		err = city.SaveFromAPI(context.Background(), "", "", elem[0])
		if err != nil {
			log.Fatal("initCities - city.SaveFromAPI", err)
		}
	}
}
