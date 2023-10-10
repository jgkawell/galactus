package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jgkawell/galactus/pkg/chassis"
	ct "github.com/jgkawell/galactus/pkg/chassis/context"

	domain1 "github.com/jgkawell/galactus/api/gen/go/domain1/aggregate1/v1"

	"github.com/google/uuid"
)

func main() {
	b := chassis.NewMainBuilder(&chassis.MainBuilderConfig{
		ApplicationName: "publisher",
		EventConfig:     &chassis.EventConfig{},
		OnRun: func(b chassis.MainBuilder) {
			ctx, _ := ct.NewExecutionContext(context.Background(), b.GetLogger(), uuid.NewString())
			manager := b.GetEventManager()
			// publish a new message every second
			for {
				id := uuid.New().String()
				event := &domain1.Created{
					Id: id,
				}
				err := manager.CreateAndSendEvent(ctx, event)
				if err != nil {
					b.GetLogger().WithError(err).Fatal("failed to emit event")
				}
				fmt.Printf("Published: %s\n", id)
				time.Sleep(1 * time.Second)
			}
		},
	})
	defer b.Close()
	b.Run()
}
