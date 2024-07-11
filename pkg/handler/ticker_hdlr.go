package handler

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type TickerHandler struct {
	ticker   *time.Ticker
	doneChan chan bool
	job      IJob
	runFirst bool
}

func NewTickerHandler(job IJob, sleep time.Duration, runFirst bool) *TickerHandler {
	doneChan := make(chan bool)
	ticker := time.NewTicker(sleep)

	return &TickerHandler{
		doneChan: doneChan,
		ticker:   ticker,
		job:      job,
		runFirst: runFirst,
	}
}

func (t *TickerHandler) error(err error, method string, params ...interface{}) error {
	return fmt.Errorf("TickerHandler.(%v)(%v) %w", method, params, err)
}

func (h *TickerHandler) Start() error {
	//create a ticker that would run retrieve every time tick
	go func() {
		//run job without waiting for ticker first

		if h.runFirst {
			err := h.job.Run()
			if err != nil {
				logrus.Error(h.error(err, "Start").Error())
			}
		}

		for {
			select {
			case <-h.doneChan:
				return
			case <-h.ticker.C:
				err := h.job.Run()

				if err != nil {
					logrus.Error(h.error(err, "Start").Error())
				}
			}
		}
	}()
	return nil
}

func (h *TickerHandler) Stop() {
	h.doneChan <- true
	close(h.doneChan)
}
