package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jgkawell/galactus/pkg/chassis"
	"github.com/jgkawell/galactus/pkg/chassis/events"
	"github.com/jgkawell/galactus/pkg/chassis/messagebus"
	"github.com/jgkawell/galactus/pkg/messaging/nats"

	domain1 "github.com/jgkawell/galactus/api/gen/go/domain1/aggregate1/v1"
)

func main() {
	bus := nats.New("")
	// bus := amqp.New("")
	b := chassis.NewMainBuilder(&chassis.MainBuilderConfig{
		ApplicationName: "publisher",
		MessageBusConfig: &chassis.MessageBusConfig{
			Buses: []messagebus.Client{
				bus,
			},
		},
		OnRun: func(b chassis.MainBuilder) {
			ctx := context.Background()
			// publish a new message every second
			for {
				id := uuid.New().String()
				err := events.Emit(ctx, []messagebus.Client{
					bus,
				}, &domain1.Created{
					Id: id,
				}, false)
				if err != nil {
					b.GetLogger().WithError(err).Fatal("failed to publish event")
				}
				fmt.Printf("Published: %s\n", id)
				time.Sleep(1 * time.Second)
			}
		},
	})
	defer b.Close()
	b.Run()
}
