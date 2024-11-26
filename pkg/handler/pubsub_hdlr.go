package handler

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/kgmedia-data/gaia/pkg/msg"

	"cloud.google.com/go/pubsub"
	"github.com/sirupsen/logrus"
)

type PubMessages struct {
	pubMessages []*pubsub.Message
	mu          sync.Mutex
	createdAt   time.Time
}

func NewPubMessages() PubMessages {
	return PubMessages{
		pubMessages: []*pubsub.Message{},
		createdAt:   time.Now(),
	}
}

func (m *PubMessages) Add(msg *pubsub.Message) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.pubMessages = append(m.pubMessages, msg)
}

func (m *PubMessages) Flush() []*pubsub.Message {
	m.mu.Lock()
	defer m.mu.Unlock()
	msgs := m.pubMessages
	m.pubMessages = []*pubsub.Message{}
	m.createdAt = time.Now()
	return msgs
}

func (m *PubMessages) Len() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.pubMessages)
}

func (m *PubMessages) isTimeout(timeout time.Duration) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return time.Since(m.createdAt) > timeout
}

type PubsubHandler[T any] struct {
	proc         IProcessor[T]
	batchProc    IBatchProcessor[T]
	batchSize    int
	batchMsgs    PubMessages
	batchTimeout time.Duration
	batchTicker  *time.Ticker
	client       *pubsub.Client
	subscription *pubsub.Subscription
	coder        msg.ICoder[T]
	ctx          context.Context
}

func NewPubsubHandler[T any](subscriptionId string, projectId string,
	proc IProcessor[T]) (*PubsubHandler[T], error) {

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		return nil, err
	}

	coder := msg.JsonCoder[T](1)
	subscription := client.Subscription(subscriptionId)

	return &PubsubHandler[T]{
		proc:         proc,
		client:       client,
		subscription: subscription,
		coder:        coder,
		ctx:          ctx,
	}, nil
}

func NewPubsubHandlerWithMaxConcurrent[T any](subscriptionId string, projectId string,
	proc IProcessor[T], maxConcurrent int) (*PubsubHandler[T], error) {

	handler, err := NewPubsubHandler(subscriptionId, projectId, proc)
	if err != nil {
		return nil, err
	}

	handler.subscription.ReceiveSettings.MaxOutstandingMessages = maxConcurrent
	return handler, nil
}

func NewPubsubHandlerBatchWithMaxConcurrent[T any](subscriptionId string, projectId string,
	batchProc IBatchProcessor[T], batchSize int, batchTimeout time.Duration, maxConcurrent int) (*PubsubHandler[T], error) {

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		return nil, err
	}

	coder := msg.JsonCoder[T](1)
	subscription := client.Subscription(subscriptionId)
	ticker := time.NewTicker(batchTimeout)

	return &PubsubHandler[T]{
		batchProc:    batchProc,
		batchSize:    batchSize,
		batchMsgs:    NewPubMessages(),
		batchTimeout: batchTimeout,
		batchTicker:  ticker,
		client:       client,
		subscription: subscription,
		coder:        coder,
		ctx:          ctx,
	}, nil
}

func (t *PubsubHandler[T]) error(err error, method string, params ...interface{}) error {
	return fmt.Errorf("PubsubHandler.(%v)(%v) %w", method, params, err)
}

func (c *PubsubHandler[T]) Start() error {
	go func(i int) {
		err := c.subscription.Receive(c.ctx,
			func(ctx context.Context, pmsg *pubsub.Message) {
				if c.isBatch() {
					c.batchMsgs.Add(pmsg)
					if c.batchMsgs.Len() >= c.batchSize || c.batchMsgs.isTimeout(c.batchTimeout) {
						c.processBatch()
					}
				} else {
					// Handle received message
					m, err := c.coder.Decode(pmsg.Data)
					if err != nil {
						logrus.Errorln(c.error(err, "Start"))
					}
					recievedData := msg.Message[T]{
						Data:      m,
						Attribute: pmsg.Attributes,
					}
					errExec := c.proc.Execute(recievedData)

					if errExec != nil {
						// Send not acknowledge if exec fail
						pmsg.Nack()
						logrus.Errorln(c.error(errExec, "Start"))
					} else {
						// Acknowledge the message to remove it from the subscription if no error
						pmsg.Ack()
					}
				}

			})
		if err != nil {
			logrus.Errorln(c.error(err, "Start"))
		}
	}(1)

	// if batch process is enabled, check if the buffer timeout is reached
	if c.isBatch() {
		go func() {
			for range c.batchTicker.C {
				if c.batchMsgs.Len() > 0 && c.batchMsgs.isTimeout(c.batchTimeout) {
					logrus.Infoln("Timeout reached, processing batch")
					c.processBatch()
				}
			}
		}()
	}

	return nil
}

func (c *PubsubHandler[T]) processBatch() {
	var msgs []msg.Message[T]
	pubMsgs := c.batchMsgs.Flush()

	validIndex := []int{}

	for i, pubMsg := range pubMsgs {
		m, err := c.coder.Decode(pubMsg.Data)
		if err != nil {
			pubMsgs[i].Nack()
			logrus.Errorln(c.error(err, "Start"))
		} else {
			msgs = append(msgs, msg.Message[T]{
				Data:      m,
				Attribute: pubMsg.Attributes,
			})
			validIndex = append(validIndex, i)
		}
	}

	if len(msgs) > 0 {
		err := c.batchProc.ExecuteBatch(msgs)
		if err != nil {
			logrus.Errorln(c.error(err, "Start"))
		}
		for _, i := range validIndex {
			if err != nil {
				pubMsgs[i].Nack()
			} else {
				pubMsgs[i].Ack()
			}
		}
	}
}

func (c *PubsubHandler[T]) isBatch() bool {
	return c.batchProc != nil
}

func (c *PubsubHandler[T]) Stop() {
	defer c.client.Close()
}
