package broker

import "context"

type Broker interface {
	CreateTopic(ctx context.Context, topic string) error
	Publish(ctx context.Context, topic string, message string) error
}
