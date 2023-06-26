package mqtt

import (
	"context"
	"fmt"
	"github.com/biosvos/parseable-pipe/internal/flow/broker"
	"github.com/pkg/errors"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var _ broker.Broker = &Mqtt{}

type Mqtt struct {
	client MQTT.Client
}

func NewMqtt() (*Mqtt, error) {
	options := MQTT.NewClientOptions()
	options.AddBroker(fmt.Sprintf("ws://%v:%v", "127.0.0.1", "9001"))

	client := MQTT.NewClient(options)
	token := client.Connect()
	token.Wait()
	if token.Error() != nil {
		return nil, errors.WithStack(token.Error())
	}
	return &Mqtt{
		client: client,
	}, nil
}

func (m *Mqtt) CreateTopic(ctx context.Context, topic string) error {
	_ = ctx
	_ = topic
	return nil
}

func (m *Mqtt) Publish(ctx context.Context, topic string, message string) error {
	_ = ctx

	token := m.client.Publish(topic, 0, false, message)
	token.Wait()
	if token.Error() != nil {
		return errors.WithStack(token.Error())
	}
	return nil
}
