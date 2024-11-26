package handler

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/kgmedia-data/gaia/pkg/msg"
	"github.com/kgmedia-data/gaia/pkg/pub"
	"github.com/stretchr/testify/assert"
)

type testBatchProc int

func (testBatchProc) ExecuteBatch(messages []msg.Message[int]) error {
	data := make([]int, 0, len(messages))
	isError := false
	for _, m := range messages {
		if m.Data == 4 || m.Data == 13 {
			isError = true
		}
		data = append(data, m.Data)
	}
	if isError {
		return fmt.Errorf("error batch data: %v", data)
	} else {
		log.Printf("batch data: %v", data)
	}
	return nil
}

func TestPubsubBatchHand(t *testing.T) {
	// set environment variable
	os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8681")

	proc := testBatchProc(0)
	pubHdlr, err := NewPubsubHandlerBatchWithMaxConcurrent[int]("gaia-dev-topic-sub",
		"gaia-dev", proc, 3, time.Second*2, 2)
	assert.NoError(t, err)

	err = pubHdlr.Start()
	assert.NoError(t, err)

	pubPub, err := pub.NewPubsubPublisher[int]("gaia-dev-topic", "gaia-dev")
	assert.NoError(t, err)

	for i := 0; i < 22; i++ {
		err = pubPub.Publish(msg.Message[int]{Data: i})
		assert.NoError(t, err)
	}

	time.Sleep(time.Second * 5)

	for i := 22; i < 44; i++ {
		err = pubPub.Publish(msg.Message[int]{Data: i})
		assert.NoError(t, err)
	}

	time.Sleep(time.Second * 5)

	pubHdlr.Stop()
}
