package main

import (
	"context"
	"fmt"

	"github.com/jgkawell/galactus/pkg/chassis"
	"github.com/jgkawell/galactus/pkg/chassis/events"
	"github.com/jgkawell/galactus/pkg/chassis/messagebus"
	"github.com/jgkawell/galactus/pkg/messaging/nats"

	eventspb "github.com/cloudevents/sdk-go/binding/format/protobuf/v2/pb"
	domain1 "github.com/jgkawell/galactus/api/gen/go/domain1/aggregate1/v1"
)

func main() {
	b := chassis.NewMainBuilder(&chassis.MainBuilderConfig{
		ApplicationName: "subscriber",
		MessageBusConfig: &chassis.MessageBusConfig{
			Buses: []messagebus.Client{
				// amqp.New(""),
				nats.New(""),
			},
		},
		HandlerLayerConfig: &chassis.HandlerLayerConfig{
			CreateConsumers: func(b chassis.MainBuilder) []chassis.ConsumerConfig {
				return []chassis.ConsumerConfig{
					{
						Event: &domain1.Created{},
						Consumer: &consumer{},
						Duplicate: true,
					},
				}
			},
		},
	})
	defer b.Close()
	b.Run()
}

type consumer struct{}

func (c *consumer) Consume(ctx context.Context, event *eventspb.CloudEvent) error {
	e := &domain1.Created{}
	err := events.Extract(event, e)
	if err != nil {
		return err
	}
	fmt.Printf("Received: %s\n", e.Id)
	return nil
}
