package pub

import (
	"context"
	"fmt"

	"github.com/kgmedia-data/gaia/pkg/msg"

	"cloud.google.com/go/pubsub"
)

type PubsubPublisher[T any] struct {
	client *pubsub.Client
	topic  *pubsub.Topic
	ctx    context.Context
	coder  msg.JsonCoder[T]
}

func NewPubsubPublisher[T any](topic string, projectId string) (*PubsubPublisher[T], error) {

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		return nil, err
	}

	t := client.Topic(topic)
	coder := msg.JsonCoder[T](1)
	return &PubsubPublisher[T]{
		coder:  coder,
		client: client,
		topic:  t,
		ctx:    ctx,
	}, nil
}

func (t *PubsubPublisher[T]) error(err error, method string, params ...interface{}) error {
	return fmt.Errorf("PubsubPublisher[T].(%v)(%v) %w", method, params, err)
}

func (c *PubsubPublisher[T]) Publish(message msg.Message[T]) error {
	data, err := c.coder.Encode(message.Data)
	if err != nil {
		return c.error(err, "Publish", message.Data)
	}
	result := c.topic.Publish(c.ctx, &pubsub.Message{
		Data:       data,
		Attributes: message.Attribute,
	})

	_, err = result.Get(c.ctx)
	if err != nil {
		return c.error(err, "Publish", message.Data)
	}
	return nil
}

func (c *PubsubPublisher[T]) Close() error {
	// Close connection
	defer c.client.Close()
	return nil
}
