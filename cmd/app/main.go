package main

import (
	"github.com/Vitaly-Baidin/weather-api/config"
	"github.com/Vitaly-Baidin/weather-api/internal/app"
	"log"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
