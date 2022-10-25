package app

import (
	"context"
	"fmt"
	v1 "github.com/Vitaly-Baidin/weather-api/internal/controller/http/v1"
	"github.com/Vitaly-Baidin/weather-api/internal/facade"
	"github.com/Vitaly-Baidin/weather-api/internal/service"
	"github.com/Vitaly-Baidin/weather-api/internal/service/repo"
	"github.com/Vitaly-Baidin/weather-api/internal/service/webapi"
	"github.com/Vitaly-Baidin/weather-api/pkg/postgres"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Run() {
	db, err := postgres.New("postgres://root:rootroot@localhost:5432/weather")
	if err != nil {
		log.Fatalln(err)
	}

	// Register Repositories
	cityRepo := repo.NewCity(db)
	tempRepo := repo.NewTemperature(db)

	// Register WebAPIs
	cityAPI := webapi.NewCity()
	tempAPI := webapi.NewTemperature()

	// Register Services
	cityService := service.NewCity(cityRepo, cityAPI)
	tempService := service.NewTemperature(tempRepo, tempAPI)

	// Register Facades
	cityFacade := facade.NewCity(cityService, tempService)
	tempFacade := facade.NewTemperature(cityService, tempService)
	rv1 := v1.NewRouter(cityFacade, tempFacade)

	// create a new server
	s := http.Server{
		Addr:         ":8080",           // configure the bind address
		Handler:      rv1,               // set the default handler
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		fmt.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			fmt.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	//gracefully shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Got signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
