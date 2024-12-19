package handler

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/sirupsen/logrus"
)

type CronHandler struct {
	cron   string
	job    IJob
	gocron *gocron.Scheduler
}

func NewCronHandler(cron string, location string, job IJob) (*CronHandler, error) {
	local, err := time.LoadLocation(location)
	if err != nil {
		local = time.UTC
	}
	gocron := gocron.NewScheduler(local)
	return &CronHandler{
		cron:   cron,
		job:    job,
		gocron: gocron,
	}, nil
}

func (t *CronHandler) error(err error, method string, params ...interface{}) error {
	return fmt.Errorf("CronHandler.(%v)(%v) %w", method, params, err)
}

func (s *CronHandler) Start() error {
	res, err := s.gocron.Cron(s.cron).Do(s.job.Run)
	s.gocron.StartAsync()

	if res.IsRunning() {
		logrus.Info("Start running Scheduler")
	}
	if err != nil {
		return s.error(err, "Start")
	}
	return nil
}

func (s *CronHandler) Stop() {
	s.gocron.Stop()
	logrus.Info("Stop running Scheduler")
}
