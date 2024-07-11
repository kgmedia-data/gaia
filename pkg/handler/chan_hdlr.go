package handler

import (
	"fmt"
	"sync"

	"github.com/kgmedia-data/gaia/pkg/msg"

	"github.com/sirupsen/logrus"
)

type ChanHandler[T any] struct {
	msgChan  chan msg.Message[T]
	doneChan chan bool
	wgC      *sync.WaitGroup
	proc     IProcessor[T]
	nWorker  int
}

func NewChanHandler[T any](msgChan chan msg.Message[T], nWorker int,
	proc IProcessor[T]) *ChanHandler[T] {

	wgC := new(sync.WaitGroup)

	return &ChanHandler[T]{
		msgChan:  msgChan,
		wgC:      wgC,
		nWorker:  nWorker,
		proc:     proc,
		doneChan: make(chan bool, nWorker),
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
					err := c.proc.Execute(m)

					if err != nil {
						logrus.Errorln(c.error(err, "Start"))
					}
				}
			}
		}(i)
	}
	return nil
}

func (c *ChanHandler[T]) Stop() {
	for i := 0; i < c.nWorker; i++ {
		c.doneChan <- true
	}
	c.wgC.Wait()
	close(c.doneChan)
}
