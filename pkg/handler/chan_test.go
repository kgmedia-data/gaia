package handler

import (
	"fmt"
	"testing"

	"github.com/kgmedia-data/gaia/pkg/msg"
)

type testProc int

func (testProc) Execute(m msg.Message[x]) error {
	fmt.Printf("data: %v attribute: %v \n", m.Data, m.Attribute)
	return nil
}

type x struct {
	Name     string
	Lastname string
}

func TestChanHandler(t *testing.T) {
	msgChan := make(chan msg.Message[x])
	proc := testProc(1)

	hdlr := NewChanHandler[x](msgChan, 3, proc)

	hdlr.Start()
	for i := 0; i < 50; i++ {
		m := msg.Message[x]{
			Data: x{
				Name: fmt.Sprintf("hello %d", i),
			},
			Attribute: map[string]string{
				"attr1": fmt.Sprintf("val%d", i),
			},
		}
		msgChan <- m
	}
	hdlr.Stop()
	close(msgChan)
}
