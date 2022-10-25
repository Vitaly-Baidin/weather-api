package cronjob

import (
	"github.com/Vitaly-Baidin/weather-api/pkg/logger"
	"github.com/robfig/cron/v3"
)

type Cron struct {
	Cron *cron.Cron
	Log  logger.Logger
}

func NewCron(log logger.Logger) *Cron {
	return &Cron{Log: log}
}

func (cj *Cron) StartCron() {
	cj.Cron = cron.New()
	cj.Cron.Start()
}

func (cj *Cron) StopCron() {
	cj.Cron.Stop()
}
