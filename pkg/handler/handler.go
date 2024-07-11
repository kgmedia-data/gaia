package handler

import (
	"os"
	"os/signal"
	"sync"

	log "github.com/sirupsen/logrus"
)

type IHandler interface {
	Start() error
	Stop()
}

type Registry struct {
	handlers  map[string]IHandler
	wg        *sync.WaitGroup
	errorChan chan handlerError
}

type handlerError struct {
	err error
	id  string
}

func NewRegistry() *Registry {
	return &Registry{
		handlers:  make(map[string]IHandler),
		wg:        new(sync.WaitGroup),
		errorChan: make(chan handlerError),
	}
}

func (r *Registry) StartAll() {
	for k, s := range r.handlers {
		r.wg.Add(1)
		r.run(k, s)
	}
	r.wait()
}

func (r *Registry) Register(id string, s IHandler) {
	r.handlers[id] = s
}

func (r *Registry) StopAll() {
	for k, s := range r.handlers {
		log.Println("Stopping handler", k)
		s.Stop()
		log.Println(k, "stopped succesfully")
	}
	r.wg.Wait()
	log.Println("All handler stopped")
}

func (r *Registry) wait() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	log.Println("waiting for signal")

	select {
	case <-signalCh:
		log.Println("interrupted")
	case err := <-r.errorChan:
		log.Errorln("fatal error for service:", err.id)
		log.Errorln(err.err)
	}
}

func (r *Registry) run(k string, s IHandler) {
	go func() {
		defer r.wg.Done()
		log.Println("Starting handler", k)
		err := s.Start()
		if err != nil {
			r.errorChan <- handlerError{
				id:  k,
				err: err,
			}
		}
		log.Println(k, "started succesfully")
	}()
}
