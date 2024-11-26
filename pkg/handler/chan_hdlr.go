package handler

import (
	"fmt"
	"sync"
	"time"

	"github.com/kgmedia-data/gaia/pkg/msg"

	"github.com/sirupsen/logrus"
)

type ChanHandler[T any] struct {
	msgChan      chan msg.Message[T]
	doneChan     chan bool
	wgC          *sync.WaitGroup
	streamProc   IProcessor[T]
	batchProc    IBatchProcessor[T]
	batchSize    int
	batchTimeout time.Duration
	batchTicker  *time.Ticker
	nWorker      int
	batchMsgs    msg.Messages[T]
}

func NewChanHandler[T any](msgChan chan msg.Message[T], nWorker int,
	proc IProcessor[T]) *ChanHandler[T] {

	wgC := new(sync.WaitGroup)

	return &ChanHandler[T]{
		msgChan:    msgChan,
		wgC:        wgC,
		nWorker:    nWorker,
		streamProc: proc,
		doneChan:   make(chan bool, nWorker),
	}
}

func NewChanBatchHandler[T any](msgChan chan msg.Message[T], nWorker int,
	proc IBatchProcessor[T], batchSize int, batchTimeout time.Duration) *ChanHandler[T] {

	wgC := new(sync.WaitGroup)

	return &ChanHandler[T]{
		msgChan:      msgChan,
		wgC:          wgC,
		nWorker:      nWorker,
		batchProc:    proc,
		batchSize:    batchSize,
		batchTimeout: batchTimeout,
		batchMsgs:    msg.NewMessages[T](),
		batchTicker:  time.NewTicker(batchTimeout),
		doneChan:     make(chan bool, nWorker),
	}
}

func (t *ChanHandler[T]) error(err error, method string, params ...interface{}) error {
	return fmt.Errorf("ChanHandler.(%v)(%v) %w", method, params, err)
}

func (c *ChanHandler[T]) Start() error {

	for i := 0; i < c.nWorker; i++ {
		c.wgC.Add(1)
		go func(i int) {
			defer c.wgC.Done()
			for {
				select {
				case <-c.doneChan:
					return
				case m := <-c.msgChan:
					if c.isBatch() {
						c.batchMsgs.Add(m)
						if c.batchMsgs.Len() >= c.batchSize || c.batchMsgs.IsTimeout(c.batchTimeout) {
							err := c.batchProc.ExecuteBatch(c.batchMsgs.Flush())
							if err != nil {
								logrus.Errorln(c.error(err, "Start"))
							}
						}
					} else {
						err := c.streamProc.Execute(m)
						if err != nil {
							logrus.Errorln(c.error(err, "Start"))
						}

					}
				}
			}
		}(i)
	}

	if c.isBatch() {
		go func() {
			for {
				select {
				case <-c.doneChan:
					return
				case <-c.batchTicker.C:
					if c.batchMsgs.Len() > 0 && c.batchMsgs.IsTimeout(c.batchTimeout) {
						err := c.batchProc.ExecuteBatch(c.batchMsgs.Flush())
						if err != nil {
							logrus.Errorln(c.error(err, "Start"))
						}
					}
				}
			}
		}()
	}
	return nil
}

func (c *ChanHandler[T]) isBatch() bool {
	return c.batchProc != nil
}

func (c *ChanHandler[T]) Stop() {
	for i := 0; i < c.nWorker; i++ {
		c.doneChan <- true
	}
	c.wgC.Wait()
	close(c.doneChan)
}
