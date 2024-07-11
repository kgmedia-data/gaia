package handler

import (
	"context"
	"fmt"

	"github.com/kgmedia-data/gaia/pkg/msg"

	"cloud.google.com/go/pubsub"
	"github.com/sirupsen/logrus"
)

type PubsubHandler[T any] struct {
	proc         IProcessor[T]
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

func (t *PubsubHandler[T]) error(err error, method string, params ...interface{}) error {
	return fmt.Errorf("PubsubHandler.(%v)(%v) %w", method, params, err)
}

func (c *PubsubHandler[T]) Start() error {
	go func(i int) {
		err := c.subscription.Receive(c.ctx,
			func(ctx context.Context, pmsg *pubsub.Message) {
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

			})
		if err != nil {
			logrus.Errorln(c.error(err, "Start"))
		}
	}(1)
	return nil
}

func (c *PubsubHandler[T]) Stop() {
	defer c.client.Close()
}
